package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	userRepository "sthannahl/usercenter/model/userrepository"

	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
)

var srv *server.Server

// InitAPIRouter .
func InitAPIRouter(oauth2Srv *server.Server) {
	srv = oauth2Srv
	http.HandleFunc("/token", tokenHandle)
	http.HandleFunc("/signUp", signUpHandle)
	http.HandleFunc("/user", userHandle)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	token, _, err := validToken(w, r)
	if err != nil {
		return
	}

	var names []string = r.URL.Query()["user_id"]
	if names == nil || len(names) == 0 || names[0] == "" {
		http.Error(w, "user_id缺失", http.StatusBadRequest)
		return
	}

	user := userRepository.GetInstance().FindUserByTypeAndName(token.GetClientID(), names[0])

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(user)
}

func signUpHandle(w http.ResponseWriter, r *http.Request) {
	token, _, err := validToken(w, r)
	if err != nil {
		return
	}

	var user map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)

	err = validUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user["type"] = token.GetClientID()
	err = userRepository.GetInstance().Save(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(&user)
}

func tokenHandle(w http.ResponseWriter, r *http.Request) {
	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validUser(user *map[string]interface{}) error {
	vaildField := []string{"user_id"}
	for _, field := range vaildField {
		if (*user)[field] == nil {
			err := errors.New("用户信息必填字段" + field + "缺失")
			return err
		}
	}
	return nil
}

func validToken(w http.ResponseWriter, r *http.Request) (oauth2.TokenInfo, oauth2.ClientInfo, error) {
	token, err := srv.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, nil, err
	}
	cli, err := srv.Manager.GetClient(token.GetClientID())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, nil, err
	}
	return token, cli, err
}
