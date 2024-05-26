package main

import (
	"wechatbot/bootstrap"
	"wechatbot/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.Run()
	gin.SetMode("debug")

	//启动web服务
	r := gin.Default()
	routers.RegisterRouter(r)

	r.Run(":" + "9527")
}
