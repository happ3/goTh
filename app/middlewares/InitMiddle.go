package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitMiddle(ctx *gin.Context) {
	fmt.Println("全局中间件开始1")
	ctx.Set("testName", "123456")
	ctx.Next()

	fmt.Println("全局中间件结束1")

}

func InitMiddle2(ctx *gin.Context) {
	fmt.Println("全局中间件开始2")

	ctx.Next()

	fmt.Println("全局中间件结束2")

}
