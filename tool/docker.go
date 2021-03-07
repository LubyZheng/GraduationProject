package tool

import (
	"fmt"
	"os/exec"
)

const (
	DOCKER = "docker"
	BUILD  = "build"
	RUN    = "run"
)

func CreateImage(name string) error {
	var args = []string{BUILD, "-t", name, "."} // docker build -t <name> .
	cmd := exec.Command(DOCKER, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Create image failed:%s\n", err.Error())
		return err
	}
	return nil
}

func CreateContainer(ImageName, TempFileName, Language string) error {
	var args = []string{RUN, ImageName, "-f", TempFileName, "-l", Language} // docker run <name> <args>
	cmd := exec.Command(DOCKER, args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Execute docker run failed:%s\n", err.Error())
		return err
	}
	fmt.Print(string(result))
	return nil
}

func CallDocker(ImageName, TempFileName, Language string) error {
	err := CreateImage(ImageName)
	if err != nil {
		return err
	}
	err = CreateContainer(ImageName, TempFileName, Language)
	if err != nil {
		return err
	}
	return nil
}
