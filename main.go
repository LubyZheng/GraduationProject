package main

import (
	"Gproject/web/handler"
	"Gproject/web/middleware"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
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
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.POST("/submit", handler.SubmitHandler)
	err := router.Run(":" + viper.GetString("server.port"))
	if err != nil {
		panic(err.Error())
	}
}
