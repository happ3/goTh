package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"goTh/app/middlewares"
	_ "goTh/docs"
	"goTh/goRedis"
	"goTh/mysqlConfig"
	"goTh/routers"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.Static("/static", "./static")

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("session", store))
	r.Use(middlewares.InitMiddle, middlewares.InitMiddle2)

	routers.RouterInit(r)

	mysqlConfig.InitMysql()
	goRedis.InitRedis()

	// Swagger UI 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
