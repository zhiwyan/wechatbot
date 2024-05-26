package service

import (
	"errors"
	"wechatbot/utils"

	"github.com/eatmoreapple/openwechat"
)

var Wechat = new(WechatService)

type WechatService struct{}

func (s *WechatService) SendTextMsg(req SendTextMsgReq) error {
	bot, ok := utils.BotMap[req.BotUserName]
	if !ok {
		return errors.New("没有这个机器人")
	}

	bReq := bot.Storage.Request
	info := bot.Storage.LoginInfo

	atText := "@" + req.AtNickName + " " + req.Content

	msg := openwechat.NewSendMessage(openwechat.MsgTypeText, atText, req.BotUserName, req.SendGroupName, "")
	wxSendMsgOption := &openwechat.ClientWebWxSendMsgOptions{
		BaseRequest: bReq,
		LoginInfo:   info,
		Message:     msg,
	}
	_, err := bot.Caller.Client.WebWxSendMsg(bot.Context(), wxSendMsgOption)
	if err != nil {
		return err
	}
	return nil
}
