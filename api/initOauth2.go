package api

import (
	omgo "sthannahl/usercenter/go-oauth2/mongo"
	"sthannahl/usercenter/go-oauth2/oauth2/manage"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3/generates"
	omanage "gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func InitOauth2Srv(JwtSignedKey, Dburi string) *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(omanage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(JwtSignedKey), jwt.SigningMethodHS512))

	clientStore := omgo.NewClientStore(omgo.NewConfig(Dburi, "user_center"))
	manager.MapClientStorage(clientStore)
	return server.NewServer(server.NewConfig(), manager)
}
