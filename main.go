package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"slight-url/config"
	"slight-url/src/user"
)

var DB *gorm.DB

func init() {
	config.InitENV()
	DB = config.InitDB()
}

func main() {
	app := gin.Default()

	user.UserModule(app, DB)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
