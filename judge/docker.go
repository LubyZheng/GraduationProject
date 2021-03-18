package judge

import (
	"Gproject/contract"
	"encoding/json"
	"os/exec"
	"strconv"
)

const (
	DOCKER = "docker"
	BUILD  = "build"
	RUN    = "run"
)

func CreateImage(ImageName string) string {
	var args = []string{BUILD, "-t", ImageName, "."} // docker build -t <name> .
	cmd := exec.Command(DOCKER, args...)
	err := cmd.Run()
	if err != nil {
		return CodeResult.PackExecuteFailResult(err.Error())
	}
	return ""
}

func CreateContainer(ImageName, TempFilePath, BinName, Language string, Time, Memory int) string {
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
	if err != nil {
		return CodeResult.PackExecuteFailResult(err.Error())
	}
	var temp contract.ExecuteResult
	json.Unmarshal(result, &temp)
	CodeResult.ExecuteMemory = temp.ExecuteMemory
	CodeResult.ExecuteTime = temp.ExecuteTime
	CodeResult.Outcome = "true"
	CodeResult.Detail = temp.Detail
	return CodeResult.ConJson()
}

func (c *Code) CallDocker() string {
	result := CreateImage(c.StudentID)
	if result != "" {
		return result
	}
	result = CreateContainer(c.StudentID, c.TempFilePath, c.BinName, c.Language, c.Time, c.Memory)
	return result
}
