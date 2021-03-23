package main

import (
	"Gproject/cgroup"
	"Gproject/contract"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var result contract.Result

type Judge struct {
	FileName   string
	FilePath   string
	Language   string
	QuestionID string
	Time       int
	Memory     int
}

func NewJudge() *Judge {
	return &Judge{}
}

func (j *Judge) Parse() {
	flag.StringVar(&j.FileName, "file", "", "")
	flag.StringVar(&j.FileName, "f", "", "")
	flag.StringVar(&j.FilePath, "path", "", "")
	flag.StringVar(&j.FilePath, "p", "", "")
	flag.StringVar(&j.Language, "language", "", "")
	flag.StringVar(&j.Language, "l", "", "")
	flag.StringVar(&j.QuestionID, "q", "", "")
	flag.StringVar(&j.QuestionID, "question", "", "")
	flag.IntVar(&j.Time, "time", 0, "") // ms
	flag.IntVar(&j.Time, "t", 0, "")
	flag.IntVar(&j.Memory, "Memory", 0, "") //kb
	flag.IntVar(&j.Memory, "m", 0, "")
	flag.Parse()
}

func (j *Judge) Run() string {
	var result contract.ExecuteResult
	var ExecuteCmd *exec.Cmd
	switch j.Language {
	case "java":
		ExecuteCmd = exec.Command("java", []string{"-cp", j.FilePath, j.FileName}...)
	default:
		ExecuteCmd = exec.Command(fmt.Sprintf("./%s", filepath.Join(j.FilePath, j.FileName)))
	}
	//输入
	input, _ := ioutil.ReadFile(fmt.Sprintf("%s.in", filepath.Join(contract.InputDir, j.QuestionID)))
	ExecuteCmd.Stdin = strings.NewReader(string(input))

	stdoutPipe, _ := ExecuteCmd.StdoutPipe()
	stderrPipe, _ := ExecuteCmd.StderrPipe()

	LimitTimeChannel := make(chan bool, 1)
	defer close(LimitTimeChannel)

	err := cgroup.MkCgroupDir()
	if err != nil {
		return result.PackUnknownErrorResult(err.Error())
	}
	cgroup.LimitMemory(j.Memory * 1024)

	//计时
	go func() {
		select {
		case <-LimitTimeChannel:
			return

		case <-time.After(time.Duration(j.Time) * time.Millisecond): //计时还包括了SetProc等IO操作，实际用时会更短
			ExecuteCmd.Process.Kill()
			result.Status = contract.TimeOutError
			return
		}
	}()
	//执行
	err = ExecuteCmd.Start()
	if err != nil {
		return result.PackUnknownErrorResult(err.Error())
	}
	//设置cgroup
	err = cgroup.SetProc(ExecuteCmd.Process.Pid)
	if err != nil {
		return result.PackUnknownErrorResult(err.Error())
	}
	outByte, _ := ioutil.ReadAll(stdoutPipe)
	errByte, _ := ioutil.ReadAll(stderrPipe)

	err = ExecuteCmd.Wait()
	if result.Status == contract.TimeOutError {
		return result.PackTimeOutResult()
	} else {
		LimitTimeChannel <- true //通知计时关闭
	}

	if err != nil {
		if strings.Contains(string(errByte), "kill") {
			return result.PackMemoryOutErrorResult()
		} else {
			return result.PackRunTimeErrorResult(string(errByte))
		}
	}else{
		return result.PackPassResult(
			string(outByte),
			fmt.Sprintf("%.3f", float32(ExecuteCmd.ProcessState.UserTime()+ExecuteCmd.ProcessState.SystemTime())/float32(time.Millisecond)),
			fmt.Sprintf("%d", ExecuteCmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss),
		)
	}
}

func main() {
	Judge := NewJudge()
	Judge.Parse()
	result := Judge.Run()
	fmt.Println(result)
}
