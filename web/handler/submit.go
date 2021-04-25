package handler

import (
	"Gproject/web/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SubmitHandler(ctx *gin.Context) {
	s, _ := ioutil.ReadAll(ctx.Request.Body)
	var RequestBody model.Request
	json.Unmarshal(s, &RequestBody)
	//err := ctx.Bind(&RequestBody)
	//if err != nil {
	//	return
	//}
	fileName := "temp"
	switch strings.ToLower(RequestBody.Language) {
	case "c":
		fileName += ".c"
	case "c++":
		fileName += ".cpp"
	case "go":
		fileName += ".go"
	case "java":
		fileName += ".java"
	}
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
	_, err = file.WriteString(RequestBody.Code)
	if err != nil {
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	cmd := exec.Command(
		"./sb",
		"-f", filepath.Join(pwd, fileName),
		"-qid", RequestBody.QuestionID,
	)
	cmd.Dir = "./sandbox"
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var resp model.Response
	_ = json.Unmarshal(result, &resp)
	ctx.JSON(http.StatusOK, resp)
}
