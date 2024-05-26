package sdk

type ReciveWeixinArgs struct {
	Pddurl       string `json:"pddurl"`
	BotUserName  string `json:"botUserName"`
	SendUserName string `json:"sendUserName"`
	AtNickName   string `json:"atNickName"`
}

type ReciveWeixinReply struct {
	StatusCode int64       `json:"statusCode"`
	Msg        string      `json:"msg"`
	Timestamp  int64       `json:"timestamp"`
	Data       interface{} `json:"data"`
}
