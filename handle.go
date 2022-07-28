package CoralBot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

type Handle struct {
	Host string
}

// 输出错误信息
func (h Handle) error(htype string, err error) {
	fmt.Printf("error:%v:%v\n", htype, err)
}

// SendPrivateMsg 发送私聊消息
func (h Handle) SendPrivateMsg(userId string, groupId string, message string, autoEscape string) {
	fromData := make(url.Values)
	fromData.Add("user_id", userId)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_private_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("私聊消息", err)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("私聊消息", err)
		return
	}
	status := gjson.Get(string(body), "status").String()
	if status != "ok" {
		fmt.Println("发送失败")
	}

}

// SendGroupMsg 发送群聊消息
func (h Handle) SendGroupMsg(groupId string, message string, autoEscape string) {
	fromData := make(url.Values)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("群聊消息", err)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("群聊消息", err)
		return
	}
	status := gjson.Get(string(body), "status").String()
	if status != "ok" {
		fmt.Println("发送失败")
	}
}

// SendGroupForwardMsg 发送合并转发 ( 群 )
func (h Handle) SendGroupForwardMsg(groupId string, message string) {
	fromData := make(url.Values)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_forward_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("合并转发", err)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("合并转发", err)
		return
	}
	status := gjson.Get(string(body), "status").String()
	if status != "ok" {
		fmt.Println("发送失败")
	}
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
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("消息", err)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("消息", err)
		return
	}
	status := gjson.Get(string(body), "status").String()
	if status != "ok" {
		fmt.Println("发送失败")
	}
}
