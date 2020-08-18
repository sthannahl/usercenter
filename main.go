package main

import (
	"fmt"

	"sthannahl/usercenter/config"
	"sthannahl/usercenter/model"
)

type AppConfig struct {
	Dburi string `yaml:"dburi"`
}

var appConfig AppConfig

func loadConfig() {
	config.Init(&appConfig)
}

func main() {
	loadConfig()
	model.InitDB(appConfig.Dburi)

	var userRepository model.UserRepository
	userRepository.SetClient(model.DB.Mongo)

	user := userRepository.FindOneUser()
	fmt.Println(user)
}
