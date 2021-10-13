package config

import "github.com/joho/godotenv"

func InitENV(){
	err:=godotenv.Load()
	if err!= nil {
		panic(err)
	}
}