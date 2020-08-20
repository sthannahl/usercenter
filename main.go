package main

import (
	"fmt"
	"log"
	"net/http"

	"sthannahl/usercenter/config"
	"sthannahl/usercenter/model"

	omgo "sthannahl/usercenter/go-oauth2/mongo"

	"sthannahl/usercenter/go-oauth2/oauth2/manage"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3/generates"
	omanage "gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
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

func initOauthSrv() *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(omanage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(appConfig.JwtSignedKey), jwt.SigningMethodHS512))

	clientStore := omgo.NewClientStore(omgo.NewConfig(appConfig.Dburi, "user_center"))
	manager.MapClientStorage(clientStore)
	return server.NewServer(server.NewConfig(), manager)
}

func main() {
	loadConfig()
	model.InitDB(appConfig.Dburi)
	srv := initOauthSrv()

	var userRepository model.UserRepository
	userRepository.SetClient(model.DB.Mongo)

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	user := userRepository.FindOneUser()
	fmt.Println(user)

	log.Printf("Server is running at %s port.", appConfig.Port)
	log.Fatal(http.ListenAndServe(":"+appConfig.Port, nil))
}
