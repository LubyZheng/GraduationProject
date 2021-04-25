package judger

import (
	"Gproject/sandbox/constants"
	result2 "Gproject/sandbox/executor/result"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	DOCKER = "docker"
	BUILD  = "build"
	RUN    = "run"
	RM     = "rm"
	RMI    = "rmi"
)

func CreateImage(ImageName string) string {
	var args = []string{BUILD, "-t", ImageName, "."} // executor build -t <name> .
	cmd := exec.Command(DOCKER, args...)
	err := cmd.Run()
	if err != nil {
		return CodeResult.PackUnknownErrorResult(err.Error())
	}
	return ""
}

func CreateContainer(Name, TempFilePath, BinName, Language, QuestionId string, Time, Memory int) string {
	// executor run <name> <args>
	var args = []string{
		RUN,
		"--name", Name, //容器名
		"--network", "none", //禁用网络
		"--privileged", Name, //因cgroup需要root权限，此处需要给Docker添加权限
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
		return CodeResult.PackUnknownErrorResult(err.Error())
	}
	var JsonResult result2.ExecuteResult
	json.Unmarshal(result, &JsonResult)
	switch JsonResult.Status {
	case constants.TimeOutError:
		return CodeResult.PackTimeOutErrorResult()
	case constants.MemoryOutError:
		return CodeResult.PackMemoryOutErrorResult()
	case constants.RunTimeError:
		return CodeResult.PackRunTimeErrorResult(JsonResult.Detail)
	case constants.UnknownError:
		return CodeResult.PackUnknownErrorResult(JsonResult.Detail)
	}
	output, _ := ioutil.ReadFile(fmt.Sprintf("%s.out", filepath.Join(constants.OutputDir, QuestionId)))
	if strings.Compare(JsonResult.Output, string(output)) == 0 {
		return CodeResult.PackPassResult(JsonResult.ExecuteTime, JsonResult.ExecuteMemory)
	} else {
		return CodeResult.PackWrongResult(JsonResult.ExecuteTime, JsonResult.ExecuteMemory)
	}
}

func (c *Code) CallDocker() string {
	result := CreateImage(c.StudentId)
	if result != "" {
		return result
	}
	result = CreateContainer(c.StudentId, c.TempFilePath, c.BinName, c.Language, c.QuestionId, c.Time, c.Memory)
	RemoveContainerAndImage(c.StudentId)
	return result
}

func RemoveContainerAndImage(Name string) {
	cmd := exec.Command(DOCKER, RM, Name)
	cmd.Run()
	cmd = exec.Command(DOCKER, RMI, Name)
	cmd.Run()
}
