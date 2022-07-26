package CoralBot

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Handle struct {
	Host string
}

// SendPrivateMsg 发送私聊消息
func (h Handle) SendPrivateMsg(userId string, groupId string, message string, autoEscape string) {
	fromData := make(url.Values)
	fromData.Add("user_id", userId)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	url := fmt.Sprintf("http://" + h.Host + "/send_private_msg")
	http.Post(url, "application/x-www-form-urlencoded", data)
}

// SendGroupMsg 发送群聊消息
func (h Handle) SendGroupMsg(groupId string, message string, autoEscape string) {
	fromData := make(url.Values)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_msg")
	http.Post(addr, "application/x-www-form-urlencoded", data)
}

// SendGroupForwardMsg 发送合并转发 ( 群 )
func (h Handle) SendGroupForwardMsg(groupId string, message string) {
	fromData := make(url.Values)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_forward_msg")
	http.Post(addr, "application/x-www-form-urlencoded", data)
}

// SendMsg 发送消息
func (h Handle) SendMsg(userId string, groupId string, message string, autoEscape string) {
	fromData := make(url.Values)
	fromData.Add("user_id", userId)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_msg")
	http.Post(addr, "application/x-www-form-urlencoded", data)
}
