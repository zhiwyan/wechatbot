package routers

import (
	"wechatbot/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	wechat := router.Group("/v1/wechat")
	{
		wechat.POST("/sendTextMsg", controller.Wechat.SendTextMsg)
	}
}
