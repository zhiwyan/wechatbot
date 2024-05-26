package handlers

import (
	//"wechatbot/gtp"
	"log"

	"github.com/eatmoreapple/openwechat"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	//log.Printf("msg.MsgType %+v", msg.MsgType)
	//log.Printf("msg.AppInfo %+v", msg.AppInfo)
	//log.Printf("msg.AppMsgType %+v", msg.AppMsgType)
	//if msg.MsgType == openwechat.MsgTypeApp && msg.AppMsgType == openwechat.AppMsgTypeUrl {
	//	return g.ReplyAppMsg(msg)
	//}
	return nil

}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *UserMessageHandler) ReplyText(m *openwechat.Message) error {
	// 接收私聊消息
	//sender, err := msg.Sender()
	//log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	//if UserService.ClearUserSessionContext(sender.ID(), msg.Content) {
	//	_, err = msg.ReplyText("上下文已经清空了，你可以问下一个问题啦。")
	//	if err != nil {
	//		log.Printf("response user error: %v \n", err)
	//	}
	//	return nil
	//}

	//msg.FromUserName = "zhiwyan"
	// 获取上下文，向GPT发起请求
	//requestText := strings.TrimSpace(msg.Content)
	//requestText = strings.Trim(msg.Content, "\n")
	//
	//_, err := msg.ReplyText("xyhs" + requestText)

	sender, err := m.Sender()
	if err != nil {
		log.Printf("Sender msg:%v, error: %v \n", m, err)
	}
	log.Printf("Received sender %+v", sender)
	log.Printf("Received msg %+v Text Msg : %v", m, m.Content)
	log.Printf("Received msg.url %+v", m.Url)

	friends, err := sender.Self().Friends()
	log.Printf("-------- friends :%+v error: %v \n", friends, err)

	zhiwyan := friends.GetByNickName("zhiwyan")
	log.Printf("zhiwyan======%+v\n", zhiwyan)

	req := m.Bot().Storage.Request
	user, _ := m.Bot().GetCurrentUser()
	info := m.Bot().Storage.LoginInfo

	msg := openwechat.NewSendMessage(openwechat.MsgTypeText, "xyhs"+string(m.Content), user.UserName, zhiwyan.UserName, "")
	wxSendMsgOption := &openwechat.ClientWebWxSendMsgOptions{
		BaseRequest: req,
		LoginInfo:   info,
		Message:     msg,
	}
	_, err = m.Bot().Caller.Client.WebWxSendMsg(m.Bot().Context(), wxSendMsgOption)
	if err != nil {
		log.Printf("response user error: %v \n", err)
	}
	return err

	//requestText = UserService.GetUserSessionContext(sender.ID()) + requestText
	//reply, err := gtp.Completions(requestText)
	//if err != nil {
	//	log.Printf("gtp request error: %v \n", err)
	//	msg.ReplyText("机器人神了，我一会发现了就去修。")
	//	return err
	//}
	//if reply == "" {
	//	return nil
	//}
	//
	//// 设置上下文，回复用户
	//reply = strings.TrimSpace(reply)
	//reply = strings.Trim(reply, "\n")
	//UserService.SetUserSessionContext(sender.ID(), requestText, reply)
	//reply = "本消息由 chatGPT Bot回复：\n" + reply
	//_, err = msg.ReplyText(reply)
	//if err != nil {
	//	log.Printf("response user error: %v \n", err)
	//}
	//return err
}

func (g *UserMessageHandler) ReplyCard(m *openwechat.Message) error {
	//user, err := msg.Bot.GetCurrentUser()
	//if err != nil {
	//	log.Printf("GetCurrentUser error: %v \n", err)
	//}
	//
	//friends, err := user.Friends()
	//if err != nil {
	//	log.Printf("GetCurrentUser error: %v \n", err)
	//}
	//for _, friend := range friends {
	//	fmt.Printf("friend:%+v", friend)
	//}
	// 接收群消息

	//m.FromUserName = "zhiwyan"
	sender, err := m.Sender()
	if err != nil {
		log.Printf("Sender msg:%v, error: %v \n", m, err)
	}
	log.Printf("Received sender %+v", sender)
	log.Printf("Received msg %+v Text Msg : %v", m, m.Content)
	log.Printf("Received msg.url %+v", m.Url)

	friends, err := sender.Self().Friends()
	log.Printf("-------- friends :%+v error: %v \n", friends, err)

	zhiwyan := friends.GetByNickName("zhiwyan")
	log.Printf("zhiwyan======%+v\n", zhiwyan)

	req := m.Bot().Storage.Request
	user, _ := m.Bot().GetCurrentUser()
	info := m.Bot().Storage.LoginInfo

	msg := openwechat.NewSendMessage(openwechat.AppMessageTypeUrl, string(m.Content), user.UserName, zhiwyan.UserName, "")
	wxSendMsgOption := &openwechat.ClientWebWxSendMsgOptions{
		BaseRequest: req,
		LoginInfo:   info,
		Message:     msg,
	}
	_, err = m.Bot().Caller.Client.WebWxSendAppMsgUrl(m.Bot().Context(), wxSendMsgOption)

	return nil
}

func (g *UserMessageHandler) ReplyAppMsg(m *openwechat.Message) error {

	return nil
}
