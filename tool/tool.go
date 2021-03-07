package tool

import (
	"Gproject/tool/flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Code struct {
	FilePath string
	Language string
	Time     int
	Memory   int
}

func New(d *flag.Default) *Code {
	return &Code{
		FilePath: d.FilePath,
		Language: d.Language,
		Time:     d.Time,
		Memory:   d.Memory,
	}
}

func CheckFileExist(path string) (string, error) {
	file, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	return file.Name(), err
}

func CreateTempDir() (string, error) {
	return ioutil.TempDir(".", "")
}

func CreateTempFile(name string) (*os.File, error) {
	return ioutil.TempFile(".", "")
}

func (c *Code) Run() error {
	//假设本地有待执行代码
	name, err := CheckFileExist(c.FilePath)
	if err != nil {
		return err
	}
	//读源码
	b, err := ioutil.ReadFile(c.FilePath)
	if err != nil {
		return err
	}
	//创建临时文件，放入当前路径下的code文件夹内
	name = fmt.Sprintf("%d%s", time.Now().UnixNano(), name)
	file, err := os.Create(fmt.Sprintf("./%s/%s", "code", name))
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
	//写入临时文件
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	err = CallDocker("2017322041", name, c.Language)
	if err != nil {
		return err
	}
	return nil
}
