package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type Judge struct {
	FileName string
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
	flag.StringVar(&j.Language, "language", "", "")
	flag.StringVar(&j.Language, "l", "", "")
	flag.IntVar(&j.Time, "Time", 0, "")
	flag.IntVar(&j.Time, "t", 0, "")
	flag.IntVar(&j.Memory, "m", 0, "")
	flag.IntVar(&j.Memory, "Memory", 0, "")
	flag.Parse()
}

func (j *Judge) Run() error {
	//编译
	var CompileCmd *exec.Cmd
	binName := strings.Split(j.FileName, ".")[0]
	switch j.Language {
	case "c":
		CompileCmd = exec.Command("gcc", []string{"-o", binName, j.FileName, "-lm"}...) //-lm for math.h
	case "cpp":
		CompileCmd = exec.Command("g++", []string{"-o", binName, j.FileName}...)
	case "go":
		CompileCmd = exec.Command("go", []string{"build", "-o", binName, j.FileName}...)
	case "java":
		CompileCmd = exec.Command("javac", []string{j.FileName}...)
	default:
		return fmt.Errorf("the language is not supported")
	}
	n, err := CompileCmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(n))
		return err
	}
	//运行
	var ExecuteCmd *exec.Cmd
	switch j.Language {
	case "java":
		ExecuteCmd = exec.Command("java", binName)
	default:
		ExecuteCmd = exec.Command(fmt.Sprintf("./%s", binName))
	}
	result, err := ExecuteCmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print("输出:", string(result))
	fmt.Println("编译时间:", CompileCmd.ProcessState.UserTime()+CompileCmd.ProcessState.SystemTime(),
		"编译内存:", fmt.Sprintf("%dkb", CompileCmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss))
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
