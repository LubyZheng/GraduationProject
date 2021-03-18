package judge

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func CompileCCmd(bin, src, tempPath string) *exec.Cmd {
	return exec.Command(
		"gcc",
		"-o",
		filepath.Join(tempPath, bin),
		filepath.Join(tempPath, src),
		"-lm") //-lm for math.h
}

func CompileCppCmd(bin, src, tempPath string) *exec.Cmd {
	return exec.Command(
		"g++",
		"-o",
		filepath.Join(tempPath, bin),
		filepath.Join(tempPath, src))
}

func CompileGoCmd(bin, src, tempPath string) *exec.Cmd {
	return exec.Command(
		"go",
		"build",
		"-o",
		filepath.Join(tempPath, bin),
		filepath.Join(tempPath, src))
}

func CompileJavaCmd(src, tempPath string) *exec.Cmd {
	return exec.Command(
		"javac",
		filepath.Join(tempPath, src))
}

func (c *Code) Compile() (string, error) {
	var CompileCmd *exec.Cmd
	switch c.Language {
	case "c":
		CompileCmd = CompileCCmd(c.BinName, c.FileName, c.TempFilePath)
	case "cpp":
		CompileCmd = CompileCppCmd(c.BinName, c.FileName, c.TempFilePath)
	case "go":
		CompileCmd = CompileGoCmd(c.BinName, c.FileName, c.TempFilePath)
	case "java":
		CompileCmd = CompileJavaCmd(c.FileName, c.TempFilePath)
	default:
		return "", fmt.Errorf("the language is not supported")
	}
	err := CompileCmd.Run()
	if err != nil {
		return CodeResult.PackCompileFailResult(), err
	}
	CodeResult.CompileTime = fmt.Sprintf("%.3f", float32(CompileCmd.ProcessState.UserTime()+CompileCmd.ProcessState.SystemTime())/float32(time.Millisecond))
	CodeResult.CompileMemory = fmt.Sprintf("%d", CompileCmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss)
	return "", nil
}
