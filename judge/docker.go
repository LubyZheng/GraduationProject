package judge

import (
	"fmt"
	"os/exec"
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

func CreateContainer(ImageName, TempFileName string) error {
	var args = []string{RUN, ImageName, "-f", TempFileName} // docker run <name> <args>
	cmd := exec.Command(DOCKER, args...)
	result, err := cmd.CombinedOutput()
	fmt.Print(string(result))
	if err != nil {
		return fmt.Errorf("Execute docker run failed:%s\n", err.Error())
	}
	return nil
}

func CallDocker(ImageName, TempFileName string) error {
	err := CreateImage(ImageName)
	if err != nil {
		return err
	}
	err = CreateContainer(ImageName, TempFileName)
	if err != nil {
		return err
	}
	return nil
}
