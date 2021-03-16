package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
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
	flag.IntVar(&j.Time, "Time", 0, "")
	flag.IntVar(&j.Time, "t", 0, "")
	flag.IntVar(&j.Memory, "m", 0, "")
	flag.IntVar(&j.Memory, "Memory", 0, "")
	flag.Parse()
}

func (j *Judge) Run() error {
	//运行
	var ExecuteCmd *exec.Cmd
	switch j.Language {
	case "java":
		ExecuteCmd = exec.Command("java", []string{"-cp", j.FilePath, j.FileName}...)
	default:
		ExecuteCmd = exec.Command(fmt.Sprintf("./%s", fmt.Sprintf("%s/%s", j.FilePath, j.FileName)))
	}
	result, err := ExecuteCmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print("输出:", string(result))
	fmt.Println("运行时间:", ExecuteCmd.ProcessState.UserTime()+ExecuteCmd.ProcessState.SystemTime(),
		"运行内存:", fmt.Sprintf("%dkb", ExecuteCmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
	return nil
}

func main() {
	judge := NewJudge()
	judge.Parse()
	err := judge.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
