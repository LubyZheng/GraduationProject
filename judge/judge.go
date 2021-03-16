package judge

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const CODE_DIR = "./code"

type Code struct {
	FileName     string //文件名
	TempFilePath string //临时文件夹名称
	FilePath     string //文件路径
	Language     string //编程语言，web需要,本地文件直接读文件类型
	BinName      string //可执行文件名
	Time         int    //限制时间
	Memory       int    //限制内存
	StudentID    string //学号
	QuestionId   string //题号
}

type Result struct {
	CompileTime   int64 `json:"编译时间(ms)"`
	CompileMemory int64 `json:"编译内存(kb)"`
}

var CodeResult = new(Result)

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
	return &Code{
		FileName:     fileName,
		TempFilePath: tempFilePath,
		FilePath:     a.FilePath,
		Language:     fileType,
		BinName:      binName,
		Time:         a.Time,
		Memory:       a.Memory,
		StudentID:    a.StudentID,
	}
}

func (c *Code) PrepareFile() error {
	//读目标文件
	b, err := ioutil.ReadFile(c.FilePath)
	if err != nil {
		return err
	}
	//创建临时文件夹
	c.TempFilePath, err = ioutil.TempDir(CODE_DIR, c.TempFilePath)
	if err != nil {
		return err
	}
	//创建临时文件
	file, err := os.Create(fmt.Sprintf("%s/%s", c.TempFilePath, c.FileName))
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

func (c *Code) Run() error {
	err := c.PrepareFile()
	if err != nil {
		return err
	}
	defer os.RemoveAll(c.TempFilePath)
	err = c.Compile()
	if err != nil {
		return err
	}
	err = c.CallDocker()
	if err != nil {
		return err
	}
	return nil
}

func Judge() {
	f := NewFlag()
	f.Parse(os.Args[1:])
	code := NewCode(f)
	err := code.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(CodeResult)
}
