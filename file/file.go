package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	MakefilePath   = "./template/Makefile"
	DockerfilePath = "./template/Dockerfile"
	AnswerPath     = "./answer/answer"
	ShellFileName  = "makeImage.sh"
	DEMO           = "demo"
)

var wg sync.WaitGroup

func CheckFileExist(path string) (string, error) {
	file, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	return file.Name(), err
}

func MakeDir() (string, error) {
	str := strconv.Itoa(int(time.Now().UnixNano()))
	err := os.Mkdir(str, 0755)
	return str, err
}

func NewMakefile(parentPath, name string) error {
	_, err := CheckFileExist(MakefilePath)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(MakefilePath)
	if err != nil {
		return err
	}
	target := fmt.Sprintf("\n\ntarget=%s", name)
	b = append(b, []byte(target)...)
	file, err := os.Create(fmt.Sprintf("./%s/Makefile", parentPath))
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func NewDockerfile(parentPath string) error {
	_, err := CheckFileExist(DockerfilePath)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(DockerfilePath)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("./%s/Dockerfile", parentPath))
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func NewTestFile(parentPath, filePath string) error {
	name, err := CheckFileExist(filePath)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("./%s/%s", parentPath, name))
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func CreateEnv(parentPath, name, filePath string) {
	go func() {
		wg.Add(1)
		NewMakefile(parentPath, name)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		NewDockerfile(parentPath)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		NewTestFile(parentPath, filePath)
		wg.Done()
	}()
	wg.Wait()
}

func CheckAnswer(output string) string {
	b, _ := ioutil.ReadFile(AnswerPath)
	switch output == string(b) {
	case true:
		return "答案正确"
	default:
		return "答案错误"
	}
}
