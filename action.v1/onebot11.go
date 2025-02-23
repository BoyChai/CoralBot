package actionv1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	botv1 "github.com/BoyChai/CoralBot/bot.v1"
	"github.com/BoyChai/CoralBot/structure"
)

type Onebot11Action struct {
	Host      string
	Agreement string
}

// 快速回复
func (action Onebot11Action) Reply(e botv1.Onebot11Event, m structure.QQMsg) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch e.MessageType {
	case "group":
		m.GroupId = e.GroupID
		data, err = action.SendMsg(e.MessageType, m)
	case "private":
		m.UserId = e.UserID
		data, err = action.SendMsg(e.MessageType, m)
	}
	return data, err

}

// 发送消息
func (action Onebot11Action) SendMsg(msgType string, m structure.QQMsg) (map[string]interface{}, error) {
	var data map[string]interface{}
	messageData := map[string]interface{}{
		"message_type": msgType,
		"user_id":      m.UserId,
		"group_id":     m.GroupId,
		"message":      m.Message,
	}
	jsonData, err := json.Marshal(messageData)
	if err != nil {
		return data, err
	}

	addr := fmt.Sprintf("%s://%s/send_msg", action.Agreement, action.Host)

	request, err := http.Post(addr, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return data, err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
