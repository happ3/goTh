package controllers

import (
	"goTh/app/models"
	"goTh/app/models/response"
	"goTh/app/models/reuqest"
	"goTh/app/service"
	"goTh/goJwt"
	"goTh/goRedis"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const jwtSecret = "ADjhjkdsfSFjhdskjf2432rf"

type UserController struct {
}

func (UserController) GetUserInfo(c *gin.Context) {
	idstr := c.Query("id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	user := models.User{
		Id: int(id),
	}
	service.GteUserInfo(&user)
	response.SuccessData(c, user)
}

/*
*
1，判断用户是否存在
2，判断用户名密码是否正确
3，加入到session中
*/
func (UserController) Login(c *gin.Context) {
	//user, err := reuqest.GetJsonToObj[models.User](c)
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		response.FailMsg(c, "数据解析失败")
	}
	service.GteUserInfo(&user)
	if user.Id == 0 {
		response.FailMsg(c, "用户不存在")
		return
	}
	infoFlag := service.CheckUserInfo(&user)
	if infoFlag {
		response.FailMsg(c, "用户名或密码不正确")
		return
	}
	//saveSession(c, &user)

	token, err := goJwt.CreateAccessToken(&user)
	if err != nil {
		response.FailMsg(c, err.Error())
		return
	}

	refreshToken := uuid.New().String()
	goRedis.RDB.Set(c, "refresh_token:"+refreshToken, user.Id, goJwt.RefreshTokenTTL).Err()

	c.SetCookie("refresh_token", refreshToken, int(goJwt.RefreshTokenTTL.Seconds()), "/", "", true, true)

	response.SuccessData(c, gin.H{"token": token, "refresh_token": refreshToken})
}

func saveSession(c *gin.Context, user *models.User) {
	session := sessions.Default(c)
	userinfoSlice, _ := json.Marshal(user)
	session.Set("userInfo", string(userinfoSlice))
	session.Save()
}

/*
*	退出登录
 */
func (UserController) Logout(c *gin.Context) {
	//delSession(c)
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		refreshToken = c.GetHeader("refresh_token")
	}
	if refreshToken != "" {
		goRedis.RDB.Del(c, "refresh_token:"+refreshToken)         // 从 Redis 删除
		c.SetCookie("refresh_token", "", -1, "/", "", true, true) // 清除浏览器 Cookie
	}

	// ▼ 第二步：【关键】吊销当前 Access Token（加入黑名单）
	authHeader := c.GetHeader("Authorization")
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if authHeader == "" || tokenStr == "" {
		c.Status(http.StatusNoContent)
		return
	}

	// 解析 Token 的过期时间（不验证签名）
	expTime, err := goJwt.GetTokenExp(tokenStr)
	if err != nil {
		// Token 格式无效，忽略
		c.Status(http.StatusNoContent)
		return
	}

	// 如果 Token 还没过期，加入黑名单，并设置 TTL = 剩余有效期
	now := time.Now()
	if expTime.After(now) {
		remaining := time.Until(*expTime)
		goRedis.RDB.Set(c, "blacklist:"+tokenStr, "1", remaining)
	}
	response.SuccessMsg(c, "退出成功")
}

// 刷新token
func (UserController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		refreshToken = c.GetHeader("refresh_token")
	}
	if len(refreshToken) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh_token 不存在"})
		return
	}

	userID, err := goRedis.RDB.Get(c, "refresh_token:"+refreshToken).Result()
	if err == redis.Nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refreshToken已过期"})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "读取Redis中的refresh_token 异常"})
		return
	}
	id64, err := strconv.ParseInt(userID, 10, 64)
	user := models.User{Id: int(id64)}
	service.GteUserInfo(&user)

	newAT, err := goJwt.CreateAccessToken(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "创建新的token失败"})
		return
	}

	response.SuccessData(c, gin.H{
		"access_token": newAT,
		"expires_in":   int(goJwt.AccessTokenTTL.Seconds()),
	})

}

func delSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userInfo")
	session.Save()
	response.SuccessMsg(c, "退出成功")
}

/*
*
校验 用户名，用户密码是否为空
创建 jwt 生成token
*/
func (UserController) RegisterUser(c *gin.Context) {
	user, err := reuqest.GetJsonToObj[models.User](c)
	if err != nil {
		response.FailMsg(c, "传递的数据不是json")
		return
	}
	if len(user.Username) == 0 {
		response.FailMsg(c, "用户名不能为空")
		return
	}
	if len(user.Password) == 0 {
		response.FailMsg(c, "用户密码不能为空")
		return
	}
	if len(user.Email) == 0 {
		response.FailMsg(c, "用户邮箱不能为空")
		return
	}

	service.Add(&user)
	response.SuccessMsg(c, "注册成功")
}

// 用户认证
func (UserController) Authentication(c *gin.Context) {
	user, err := reuqest.GetJsonToObj[models.User](c)
	if err != nil {
		response.FailMsg(c, "传递的数据不是json")
		return
	}
	if user.Id == 0 {
		response.FailMsg(c, "id不能为空")
		return
	}
	user.Authentication = "1"
	err = service.UpdateUserInfo(&user)
	if err != nil {
		response.FailMsg(c, err.Error())
		return
	}
	response.SuccessMsg(c, "认证成功")
}
