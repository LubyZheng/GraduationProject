package main

import (
	"Gproject/file"
	"Gproject/flags"
	"fmt"
	"os"
	"os/exec"
)

//暂时只支持Go语言
func main() {
	f := flags.New()
	f.Parse(os.Args[1:])

	//判断文件是否存在
	name, err := file.CheckFileExist(f.FilePath)
	if err != nil {
		fmt.Printf("file does not exist:%s\n", err.Error())
		return
	}

	//新建唯一目录
	path, err := file.MakeDir()
	if err != nil {
		fmt.Printf("create directory failed:%s\n", err.Error())
		return
	}
	defer func() {
		os.RemoveAll(path)
	}()

	//统一放到唯一目录下，为生成新镜像创造环境
	file.CreateEnv(path, name, f.FilePath)

	//生成新镜像 TODO：每跑一个新程序都要生成一个新镜像且镜像名固定，先完成基础功能，待改进
	cmd := exec.Command("sh", file.ShellFileName, path)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Create image failed:%s\n", err.Error())
		return
	}

	//运行docker容器，执行待检测的程序
	cmd = exec.Command("docker", []string{"run", file.DEMO}...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Execute docker run failed:%s\n", err.Error())
		return
	}

	fmt.Println(string(result))

	//对比答案
	fmt.Println(file.CheckAnswer(string(result)))
}
