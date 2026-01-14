package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"goTh/app/models"
)

func GetSession(c *gin.Context) (u models.User) {
	session := sessions.Default(c)
	jsonVal := session.Get("userInfo")
	user := models.User{}
	json.Unmarshal([]byte(jsonVal.(string)), &user)
	return user
}
