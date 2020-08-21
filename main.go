package main

import (
	"fmt"
	"log"
	"net/http"

	"sthannahl/usercenter/api"
	"sthannahl/usercenter/config"
	"sthannahl/usercenter/model"
)

type AppConfig struct {
	Dburi        string `yaml:"dburi"`
	JwtSignedKey string `yaml:"jwtSignedKey"`
	Port         string `yaml:"port"`
}

var appConfig AppConfig

func loadConfig() {
	config.Init(&appConfig)
}

func main() {
	loadConfig()
	model.InitDB(appConfig.Dburi)

	srv := api.InitOauth2Srv(appConfig.JwtSignedKey, appConfig.Dburi)
	api.InitApiRouter(srv)

	var userRepository model.UserRepository
	userRepository.SetClient(model.DB.Mongo)
	user := userRepository.FindOneUser()
	fmt.Println(user)

	log.Printf("Server is running at %s port.", appConfig.Port)
	log.Fatal(http.ListenAndServe(":"+appConfig.Port, nil))
}
