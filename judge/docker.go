package judge

import (
	"Gproject/contract"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
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

func CreateContainer(Name, TempFilePath, BinName, Language, QuestionId string, Time, Memory int) string {
	// docker run <name> <args>
	var args = []string{
		RUN,
		"--name", Name, //容器名
		"--privileged", Name, //镜像名
		"-p", TempFilePath,
		"-q", QuestionId,
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
	fmt.Println(string(result))
	var temp contract.ExecuteResult
	json.Unmarshal(result, &temp)
	output, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s.out", contract.OutputDir, QuestionId))
	if strings.Compare(temp.Output, string(output)) == 0 {
		CodeResult.Outcome = "true"
	} else {
		CodeResult.Outcome = "false"
	}
	CodeResult.ExecuteMemory = temp.ExecuteMemory
	CodeResult.ExecuteTime = temp.ExecuteTime
	CodeResult.Detail = temp.Detail
	return CodeResult.ConJson()
}

func (c *Code) CallDocker() string {
	result := CreateImage(c.StudentID)
	if result != "" {
		return result
	}
	result = CreateContainer(c.StudentID, c.TempFilePath, c.BinName, c.Language, c.QuestionId, c.Time, c.Memory)
	RemoveContainerAndImage(c.StudentID)
	return result
}

func RemoveContainerAndImage(Name string) {
	cmd := exec.Command("docker", "rm", Name)
	cmd.Run()
	cmd = exec.Command("docker", "rmi", Name)
	cmd.Run()
}
