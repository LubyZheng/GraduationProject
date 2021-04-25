package main

import (
	"Gproject/sandbox/constants"
	"Gproject/sandbox/executor/cgroup"
	result2 "Gproject/sandbox/executor/result"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var result result2.ExecuteResult

type Flag struct {
	FileName   string
	FilePath   string
	Language   string
	QuestionID string
	Time       int
	Memory     int
}

func NewFlag() *Flag {
	return &Flag{}
}

func (f *Flag) Parse() {
	flag.StringVar(&f.FileName, "file", "", "")
	flag.StringVar(&f.FileName, "f", "", "")
	flag.StringVar(&f.FilePath, "path", "", "")
	flag.StringVar(&f.FilePath, "p", "", "")
	flag.StringVar(&f.Language, "language", "", "")
	flag.StringVar(&f.Language, "l", "", "")
	flag.StringVar(&f.QuestionID, "q", "", "")
	flag.StringVar(&f.QuestionID, "question", "", "")
	flag.IntVar(&f.Time, "time", 0, "") // ms
	flag.IntVar(&f.Time, "t", 0, "")
	flag.IntVar(&f.Memory, "Memory", 0, "") //kb
	flag.IntVar(&f.Memory, "m", 0, "")
	flag.Parse()
}

type Judge struct {
	FileName   string
	FilePath   string
	Language   string
	QuestionID string
	Time       time.Duration
	Memory     int
}

func NewJudge(f *Flag) *Judge {
	return &Judge{
		FileName:   f.FileName,
		FilePath:   f.FilePath,
		Language:   f.Language,
		QuestionID: f.QuestionID,
		Time:       time.Duration(f.Time) * time.Millisecond, //ms
		Memory:     f.Memory * 1024,                          //kb
	}
}

func (j *Judge) Run() string {
	var ExecuteCmd *exec.Cmd
	switch j.Language {
	case "java":
		ExecuteCmd = exec.Command("java", []string{"-cp", j.FilePath, j.FileName}...)
	default:
		ExecuteCmd = exec.Command(fmt.Sprintf("./%s", filepath.Join(j.FilePath, j.FileName)))
	}
	//输入
	input, _ := ioutil.ReadFile(fmt.Sprintf("%s.in", filepath.Join(constants.InputDir, j.QuestionID)))
	ExecuteCmd.Stdin = strings.NewReader(string(input))

	stdoutPipe, _ := ExecuteCmd.StdoutPipe()
	stderrPipe, _ := ExecuteCmd.StderrPipe()

	LimitTimeChannel := make(chan bool, 1)
	defer close(LimitTimeChannel)

	err := cgroup.MkCgroupDir()
	if err != nil {
		return result.PackUnknownErrorResult(err.Error())
	}
	err = cgroup.LimitMemory(j.Memory)
	if err != nil {
		return result.PackUnknownErrorResult(err.Error())
	}

	//计时
	go func() {
		select {
		case <-LimitTimeChannel:
			return

		case <-time.After(j.Time): //计时还包括了SetProc等IO操作，实际用时会更短
			ExecuteCmd.Process.Kill()
			result.Status = constants.TimeOutError
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

	if result.Status == constants.TimeOutError {
		return result.PackTimeOutResult()
	} else {
		LimitTimeChannel <- true //通知计时关闭
	}
	if err != nil {
		if strings.Contains(err.Error(), "kill") {
			return result.PackMemoryOutErrorResult()
		} else {
			return result.PackRunTimeErrorResult(string(errByte))
		}
	} else {
		mem, _ := strconv.Atoi(cgroup.ReadMemory())
		mem = mem / 1024 //kb
		return result.PackPassResult(
			string(outByte),
			fmt.Sprintf("%.3f", float32(ExecuteCmd.ProcessState.UserTime()+ExecuteCmd.ProcessState.SystemTime())/float32(time.Millisecond)),
			fmt.Sprintf("%d", mem),
		)
	}
}

func main() {
	Flag := NewFlag()
	Flag.Parse()
	Judge := NewJudge(Flag)
	result := Judge.Run()
	fmt.Println(result)
}
