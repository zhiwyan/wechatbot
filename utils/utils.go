package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func GetPddGoodId(rawURL string) string {
	// 定义正则表达式
	re := regexp.MustCompile(`goods_id=([0-9]+)`)

	// 查找匹配的子字符串
	match := re.FindStringSubmatch(rawURL)
	if len(match) < 2 {
		fmt.Println("未找到goods_id参数")
		return ""
	}

	goodsID := match[1]
	fmt.Println("goods_id:", goodsID)
	return goodsID
}

func GetPddUri(link string) string {
	re := regexp.MustCompile(`https://mobile\.yangkeduo\.com/([^/?]+)\?(.*)`)
	match := re.FindStringSubmatch(link)
	if len(match) > 1 {
		page := match[1]
		params := match[2]
		reParams := regexp.MustCompile(`goods_id=[^&]+`)
		goodsIdMatch := reParams.FindString(params)
		return fmt.Sprintf("https://mobile.yangkeduo.com/%s?%s", page, goodsIdMatch)
	}
	return ""
}

func PostJson(url string, params interface{}) ([]byte, error) {
	headers := make(map[string]string, 1)

	return PostJsonHeader(url, params, headers)
}

func PostJsonHeader(url string, params interface{}, headers map[string]string) ([]byte, error) {
	//url := "http://adminliveapi.bccv5.vdyoo.com/classroom/GetRoomList"
	//payload := strings.NewReader(`{
	//	"bizId": 3,
	//	"planId": 2871329
	//}`)

	paramByte, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(paramByte)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("httpcode error:" + fmt.Sprint(res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
