package main

import (
	"fmt"
	"net/http"

	"sthannahl/usercenter/config"
	"sthannahl/usercenter/model"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type AppConfig struct {
	Dburi string `yaml:"dburi"`
}

var appConfig AppConfig

func loadConfig() {
	config.Init(&appConfig)
}

func a() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("sthannahl"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
}

func main() {
	loadConfig()
	model.InitDB(appConfig.Dburi)

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
}
