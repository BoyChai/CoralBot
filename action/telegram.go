package action

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TelegramHandle struct {
	// tgAPIKe
	Token string
}

func (t TelegramHandle) GetHandlerType() string {
	return "Telegram"
}

func (t TelegramHandle) getData(method string, data url.Values) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", t.Token, method)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.New("TG请求创建错误:" + err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 采用代理
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("TG方法%s发送失败,错误信息:%s", method, err.Error()))
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("API返回数据读取错误:" + err.Error())
	}
	return respData, nil
}
