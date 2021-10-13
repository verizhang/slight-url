package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDns() (dns string){
	username  := os.Getenv("DB_USERNAME")
	password  := os.Getenv("DB_PASSWORD")
	host  := os.Getenv("DB_HOST")
	port  := os.Getenv("DB_PORT")
	name  := os.Getenv("DB_NAME")
	dns  = username+":"+password+"@tcp("+host+":"+port+")/"+name+"?charset=utf8mb4&parseTime=True&loc=Local"
	return
}

func InitDB() (DB *gorm.DB){
	dns := GetDns()
	DB, err := gorm.Open(mysql.Open(dns))
	if err!= nil{
		panic(err)
	}
	return
}