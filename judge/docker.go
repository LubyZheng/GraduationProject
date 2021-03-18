package judge

import (
	"fmt"
	"os/exec"
	"strconv"
)

const (
	DOCKER = "docker"
	BUILD  = "build"
	RUN    = "run"
)

func CreateImage(ImageName string) error {
	var args = []string{BUILD, "-t", ImageName, "."} // docker build -t <name> .
	cmd := exec.Command(DOCKER, args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Create image failed:%s\n", err.Error())
	}
	return nil
}

func CreateContainer(ImageName, TempFilePath, BinName, Language string, Time, Memory int) error {
	// docker run <name> <args>
	var args = []string{
		RUN, ImageName,
		"-p", TempFilePath,
		"-f", BinName,
		"-l", Language,
		"-t", strconv.Itoa(Time),
		"-m", strconv.Itoa(Memory),
	}
	cmd := exec.Command(DOCKER, args...)
	result, err := cmd.CombinedOutput()
	fmt.Print(string(result))
	if err != nil {
		return fmt.Errorf("Execute docker run failed:%s\n", err.Error())
	}
	return nil
}

func (c *Code) CallDocker() error {
	err := CreateImage(c.StudentID)
	if err != nil {
		return err
	}
	err = CreateContainer(c.StudentID, c.TempFilePath, c.BinName, c.Language, c.Time, c.Memory)
	if err != nil {
		return err
	}
	return nil
}
