package middlewares

import (
	"errors"
	"goTh/app/models"
	"goTh/app/models/response"
	"goTh/goJwt"
	"goTh/goRedis"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

/*
*
1，验证 token是否有效
2，验证是否过期
3，验证是否在黑名单中
*/
func AuthMiddleware(c *gin.Context) {
	authHeaderToken := c.GetHeader("Authorization")
	if authHeaderToken == "" {
		response.FailMsg(c, "token不存在")
		c.Abort()
		return
	}

	userClaims, err := goJwt.ParseAccessToken(authHeaderToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			response.FailMsg(c, "token已过期")
		} else {
			response.FailMsg(c, "token无效")
		}
		c.Abort()
		return
	}
	user := models.User{Id: userClaims.UserId, Username: userClaims.UserName}

	// 关键：检查黑名单（即使 Token 未过期，也可能已被吊销）
	exists, _ := goRedis.RDB.Exists(c, "blacklist:"+authHeaderToken).Result()
	if exists > 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token has been revoked"})
		return
	}

	c.Set("user", &user)
	c.Next()
}
