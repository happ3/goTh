package main

import (
	"fmt"
	"goTh/goRedis"
	"goTh/mysqlConfig"
	"goTh/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.Static("/static", "./static")

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("session", store))

	routers.RouterInit(r)

	mysqlConfig.InitMysql()
	goRedis.InitRedis()

	r.Run(":8080")
}
