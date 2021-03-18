package main

import (
	"Gproject/contract"
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type Judge struct {
	FileName string
	FilePath string
	Language string
	Time     int
	Memory   int
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
	flag.IntVar(&j.Time, "time", 0, "")
	flag.IntVar(&j.Time, "t", 0, "")
	flag.IntVar(&j.Memory, "Memory", 0, "")
	flag.IntVar(&j.Memory, "m", 0, "")
	flag.Parse()
}

func (j *Judge) Run() string {
	//运行
	var result contract.ExecuteResult
	var ExecuteCmd *exec.Cmd
	switch j.Language {
	case "java":
		ExecuteCmd = exec.Command("java", []string{"-cp", j.FilePath, j.FileName}...)
	default:
		ExecuteCmd = exec.Command(fmt.Sprintf("./%s", fmt.Sprintf("%s/%s", j.FilePath, j.FileName)))
	}
	LimitTimeChannel := make(chan bool, 1)
	defer close(LimitTimeChannel)
	//计时
	go func() {
		select {
		case <-LimitTimeChannel:
			return

		case <-time.After(time.Duration(j.Time) * time.Millisecond):
			ExecuteCmd.Process.Kill()
			return
		}
	}()
	_, err := ExecuteCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(err.Error(), "kill") {
			LimitTimeChannel <- true
			return result.PackTimeOutResult()
		} else {
			return result.PackRunTimeErrorResult(err.Error())
		}
	} else {
		LimitTimeChannel <- true
		return result.PackPassResult(
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
