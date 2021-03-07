package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Judge struct {
	File     string
	Language string
	Time     int
	Memory   int
}

func NewJudge(d *Default) *Judge {
	return &Judge{
		File:     d.FilePath,
		Language: d.Language,
		Time:     d.Time,
		Memory:   d.Memory,
	}
}

type Default struct {
	flagSet  *flag.FlagSet
	FilePath string
	Language string
	Time     int
	Memory   int
}

func NewFlag() *Default {
	return &Default{
		flagSet: flag.NewFlagSet("", flag.ExitOnError),
	}
}

func (d *Default) Parse(args []string) {
	d.flagSet.StringVar(&d.FilePath, "file", "", "")
	d.flagSet.StringVar(&d.FilePath, "f", "", "")
	d.flagSet.StringVar(&d.Language, "language", "", "")
	d.flagSet.StringVar(&d.Language, "l", "", "")
	d.flagSet.IntVar(&d.Time, "Time", 0, "")
	d.flagSet.IntVar(&d.Time, "t", 0, "")
	d.flagSet.IntVar(&d.Memory, "m", 0, "")
	d.flagSet.IntVar(&d.Memory, "Memory", 0, "")
	d.flagSet.Usage = d.helpCallback
	err := d.flagSet.Parse(args)
	if err != nil {
		os.Exit(0)
	}
}

func (d *Default) helpCallback() {
	fmt.Printf(
		"Usage: %s [options]\n"+
			"Options:\n"+
			"    -f, --file <name>               FilePath to be executed. ex: xxx.cpp\n"+
			"    -l, --language <language>       Code language. ex: C++\n"+
			"    -t, --Time <Time>               Limit Time. unit: ms\n"+
			"    -m, --momery <Memory>           Limit Memory. unit: kb\n"+
			"Common Options:\n"+
			"	 -h, --help                      Show this message\n",
		os.Args[0],
	)
	os.Exit(0)
}

func (j *Judge) Run() error {
	//编译
	var CompileCmd *exec.Cmd
	switch strings.ToUpper(j.Language) {
	case "C":
		CompileCmd = exec.Command("gcc", []string{"-o", "demo", j.File}...)
	case "C++":
		CompileCmd = exec.Command("g++", []string{"-o", "demo", j.File}...)
	case "JAVA":
		CompileCmd = exec.Command("java", []string{"-o", "demo", j.File}...)
	case "GO":
		CompileCmd = exec.Command("go", []string{"build", "-o", "demo", j.File}...)
	default:
		return fmt.Errorf("the language is not supported")
	}
	err := CompileCmd.Run()
	if err != nil {
		return err
	}
	//运行
	ExecuteCmd := exec.Command("./demo")
	result, err := ExecuteCmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print("输出:", string(result))
	fmt.Println("编译时间:", CompileCmd.ProcessState.UserTime()+CompileCmd.ProcessState.SystemTime())
	fmt.Println("运行时间:", ExecuteCmd.ProcessState.UserTime()+ExecuteCmd.ProcessState.SystemTime())
	return nil
}

func main() {
	f := NewFlag()
	f.Parse(os.Args[1:])
	judge := NewJudge(f)
	err := judge.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
