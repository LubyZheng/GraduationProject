package judge

import (
	"Gproject/contract"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Code struct {
	FileName     string //文件名
	TempFilePath string //临时文件夹名称
	FilePath     string //文件路径
	Language     string //编程语言，web需要,本地文件直接读文件类型
	BinName      string //可执行文件名
	Time         int    //限制时间
	Memory       int    //限制内存
	StudentId    string //学号
	QuestionId   string //题号
}

var CodeResult = new(contract.Result)

func CheckFileExist(path string) (string, error) {
	file, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	return file.Name(), err
}

func CheckFileType(FileName string) string {
	FileNameArray := strings.Split(FileName, ".")
	FileType := FileNameArray[len(FileNameArray)-1]
	return FileType
}

func NewCode(a *Arguments) *Code {
	fileName, _ := CheckFileExist(a.FilePath)
	fileType := CheckFileType(fileName)
	tempFilePath := fmt.Sprintf("%s_%s_%s_", a.StudentID, time.Now().Format("20060102150405"), a.QuestionID)
	binName := strings.Split(fileName, ".")[0]
	if a.Time == 0 {
		a.Time = contract.DefaultExecuteTime
	}
	if a.Memory == 0 {
		a.Memory = contract.DefaultExecuteMemory
	}
	return &Code{
		FileName:     fileName,
		TempFilePath: tempFilePath,
		FilePath:     a.FilePath,
		Language:     fileType,
		BinName:      binName,
		Time:         a.Time,
		Memory:       a.Memory,
		StudentId:    a.StudentID,
		QuestionId:   a.QuestionID,
	}
}

func (c *Code) PrepareFile() error {
	//读目标文件
	b, err := ioutil.ReadFile(c.FilePath)
	if err != nil {
		return err
	}
	//创建临时文件夹
	c.TempFilePath, err = ioutil.TempDir(contract.CodeDir, c.TempFilePath)
	if err != nil {
		return err
	}
	//创建临时文件
	file, err := os.Create(filepath.Join(c.TempFilePath, c.FileName))
	if err != nil {
		return err
	}
	defer file.Close()
	//拷贝源码
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (c *Code) Run() string {
	err := c.PrepareFile()
	if err != nil {
		return err.Error()
	}
	defer os.RemoveAll(c.TempFilePath)
	result := c.Compile()
	if result != "" {
		return result
	}
	return c.CallDocker()
}

func Judge() {
	f := NewFlag()
	f.Parse(os.Args[1:])
	code := NewCode(f)
	result := code.Run()
	fmt.Println(result)
}
