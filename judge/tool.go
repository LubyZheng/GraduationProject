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

func New(a *Arguments) (*Code, error) {
	fileName, err := CheckFileExist(a.FilePath)
	if err != nil {
		return nil, err
	}
	fileType := CheckFileType(fileName)
	return &Code{
		FileName:     fileName,
		TempFilePath: fmt.Sprintf("%s_%s_%s_", a.StudentID, time.Now().Format("20060102150405"), a.QuestionID),
		FilePath:     a.FilePath,
		Language:     fileType,
		BinName:      strings.Split(fileName, ".")[0],
		Time:         a.Time,
		Memory:       a.Memory,
		StudentID:    a.StudentID,
	}, nil
}

func (c *Code) Run() error {
	//读目标文件
	b, err := ioutil.ReadFile(c.FilePath)
	if err != nil {
		return err
	}
	//创建临时文件夹
	tempDir, err := ioutil.TempDir(CODE_DIR, c.TempFilePath)
	if err != nil {
		return err
	}
	//defer os.Remove(tempDir)
	//拷贝源文件放入
	file, err := os.Create(fmt.Sprintf("%s/%s", tempDir, c.FileName))
	if err != nil {
		return err
	}
	defer file.Close()
	//写入临时文件
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	err = Compile(c.BinName, c.FileName, tempDir, c.Language)
	if err != nil {
		return err
	}
	err = CallDocker(c.StudentID, c.TempFilePath)
	if err != nil {
		return err
	}
	return nil
}
