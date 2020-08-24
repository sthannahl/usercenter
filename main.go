package main

import (
	"log"
	"net/http"

	"sthannahl/usercenter/api"
	"sthannahl/usercenter/config"
	"sthannahl/usercenter/model"
	userRepository "sthannahl/usercenter/model/userrepository"
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

	userRepository.GetInstance().SetClient(model.DB.Mongo)

	srv := api.InitOauth2Srv(appConfig.JwtSignedKey, appConfig.Dburi)
	api.InitAPIRouter(srv)

	log.Printf("Server is running at %s port.", appConfig.Port)
	log.Fatal(http.ListenAndServe(":"+appConfig.Port, nil))
	// gin.Default()
}
