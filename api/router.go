package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	userRepository "sthannahl/usercenter/model/userrepository"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
)

var srv *server.Server

// InitAPIRouter .
func InitAPIRouter(port string) {
	router := gin.Default()

	router.POST("/token", tokenHandle)
	router.POST("/signUp", signUpHandle)
	router.GET("/user", userHandle)

	log.Printf("Server is running at %s port.", port)
	log.Fatal(router.Run(":" + port))
}

// SetOauth2Srv .
func SetOauth2Srv(oauth2Srv *server.Server) {
	srv = oauth2Srv
}

func userHandle(c *gin.Context) {
	token, _, err := validToken(c)
	if err != nil {
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "user_id缺失",
		})
		return
	}

	user := userRepository.GetInstance().FindUserByTypeAndName(token.GetClientID(), userID)

	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  "",
		"data": user,
	})
}

func signUpHandle(c *gin.Context) {
	token, _, err := validToken(c)
	if err != nil {
		return
	}

	var user map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)

	err = validUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user["type"] = token.GetClientID()
	err = userRepository.GetInstance().Save(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "",
		"data": user,
	})
}

func tokenHandle(c *gin.Context) {
	err := srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
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

func validToken(c *gin.Context) (oauth2.TokenInfo, oauth2.ClientInfo, error) {
	token, err := srv.ValidationBearerToken(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return nil, nil, err
	}
	cli, err := srv.Manager.GetClient(token.GetClientID())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return nil, nil, err
	}
	return token, cli, err
}
