package main

import (
	"Gproject/web/handler"
	"Gproject/web/middleware"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//解决命令行接口调用反馈的显示问题
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	//
	router := gin.Default()
	//加载html文件
	router.LoadHTMLGlob("./templates/*")
	//加载css
	router.StaticFS("/static", http.Dir("./static"))
	//解决跨域问题
	router.Use(middleware.CORSMiddleware())
	//路由
	router.GET("/", handler.IndexHandler)
	router.GET("/problem", handler.QuestionHandler)
	router.GET("/submit", handler.SubmitGetHandler)
	router.POST("/submit", handler.SubmitPostHandler)
	err := router.Run(":" + viper.GetString("server.port"))
	if err != nil {
		panic(err.Error())
	}
}
