package controller

import (
	"fmt"
	"wechatbot/service"
	"wechatbot/utils"

	"github.com/gin-gonic/gin"
)

var Wechat = new(WechatController)

type WechatController struct {
}

func (ctl *WechatController) SendTextMsg(ctx *gin.Context) {
	var params service.SendTextMsgReq
	if err := ctx.ShouldBind(&params); err != nil {
		utils.ReturnFail(ctx, err)
		return
	}
	fmt.Printf("req %+v \v", params)

	// cdn reload
	err := service.Wechat.SendTextMsg(params)
	if err != nil {
		utils.ReturnFail(ctx, err)
		return
	}

	utils.ReturnSuccessNoCount(ctx, nil)
	return
}
