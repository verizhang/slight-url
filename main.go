package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slight-url/config"
)

func main(){
	config.InitENV()
	config.InitDB()
	app := gin.Default()

	app.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message":"success"})
	})

	err := app.Run()
	if err != nil{
		panic(err)
	}
}