package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"slight-url/config"
	"slight-url/src/modules"
)

var DB *gorm.DB

func init() {
	config.InitENV()
	config.InitDB()
}

func main() {
	app := gin.Default()

	modules.UserModule(app)
	modules.UrlModule(app)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
