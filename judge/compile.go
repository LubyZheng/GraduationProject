package judge

import (
	"fmt"
	"os/exec"
	"path/filepath"
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

func Compile(bin, src, tempPath, language string) error {
	var CompileCmd *exec.Cmd
	switch language {
	case "c":
		CompileCmd = CompileCCmd(bin, src, tempPath)
	case "cpp":
		CompileCmd = CompileCppCmd(bin, src, tempPath)
	case "go":
		CompileCmd = CompileGoCmd(bin, src, tempPath)
	case "java":
		CompileCmd = CompileJavaCmd(src, tempPath)
	default:
		return fmt.Errorf("the language is not supported")
	}
	err := CompileCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
