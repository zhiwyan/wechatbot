package sdk

import (
	"encoding/json"
	"fmt"
	"sync"
	"wechatbot/utils"
)

const HOST = "http://192.168.110.207:8080"

type Mirrorapi struct{}

var once sync.Once

var instance *Mirrorapi

func NewMirrorapi() *Mirrorapi {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		instance = new(Mirrorapi)
	})
	return instance
}

func (this *Mirrorapi) ReciveWeixin(req ReciveWeixinArgs) (ReciveWeixinReply, error) {
	resp := ReciveWeixinReply{}
	url := HOST + "/mirror/reciveWeixin"
	res, err := utils.PostJson(url, req)
	fmt.Println("-----", url, err, req, string(res))
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(res, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
