package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const failCode = 400
const successCode = 200

func FailMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": failCode,
		"msg":  msg,
	})
}
func FailData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": failCode,
		"data": data,
		"msg":  "fail",
	})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": successCode,
		"data": data,
		"msg":  "success",
	})
}
func SuccessMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": successCode,
		"msg":  msg,
	})
}
