package service

// 发送文本消息
type SendTextMsgReq struct {
	Pddurl        string `json:"pddurl" `
	BotUserName   string `json:"botUserName" binding:"required"`
	SendGroupName string `json:"sendGroupName" binding:"required"`
	AtNickName    string `json:"atNickName" binding:"required"`
	Content       string `json:"content" binding:"required"`
}
