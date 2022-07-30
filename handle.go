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

// DeleteMsg 撤回消息
func (h Handle) DeleteMsg(messageId int32) {
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/delete_msg")
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
	//该 API 无响应数据
}

// GetMsg 获取消息
func (h Handle) GetMsg(messageId int32) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("GetMsg", err)
		return e
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("GetMsg", err)
		return e
	}
	e.bodyData = string(body)
	e.MessageId = gjson.Get(e.bodyData, "data.message_id").String()
	e.RealId = gjson.Get(e.bodyData, "data.real_id").String()
	e.Sender.UserID = gjson.Get(e.bodyData, "data.sender.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "data.sender.nickname").String()
	e.Sender.Sex = gjson.Get(e.bodyData, "data.sender.sex").String()
	e.Sender.Age = gjson.Get(e.bodyData, "data.sender.age").String()
	e.Sender.Card = gjson.Get(e.bodyData, "data.sender.card").String()
	e.Sender.Area = gjson.Get(e.bodyData, "data.sender.area").String()
	e.Sender.Level = gjson.Get(e.bodyData, "data.sender.level").String()
	e.Sender.Role = gjson.Get(e.bodyData, "data.sender.role").String()
	e.Sender.Title = gjson.Get(e.bodyData, "data.sender.title").String()
	e.Time = gjson.Get(e.bodyData, "data.time").String()
	e.Message = gjson.Get(e.bodyData, "data.message").String()
	e.RawMessage = gjson.Get(e.bodyData, "data.raw_message").String()
	return e
}

// GetForwardMsg 获取合并转发内容
func (h Handle) GetForwardMsg(messageId int32) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_forward_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("GetForwardMsg", err)
		return e
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("GetForwardMsg", err)
		return e
	}
	status := gjson.Get(string(body), "status").String()
	if status != "ok" {
		fmt.Println("发送失败")
		return e
	}
	e.bodyData = string(body)
	for i := 0; i < int(gjson.Get(e.bodyData, "data.messages.#").Int()); i++ {
		var messages ForwardMessage
		messages.Content = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+string(i)+".content")).String()
		messages.Sender.Nickname = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+string(i)+".sender.nickname")).String()
		messages.Sender.UserID = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+string(i)+".sender.user_id")).String()
		messages.Time = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+string(i)+".time")).String()
		e.Messages = append(e.Messages, messages)
	}
	return e
}

// GetImage 获取图片信息
func (h Handle) GetImage(file string) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("file", file)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_image")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		h.error("get_image", err)
		return e
	}
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		h.error("get_image", err)
		return e
	}
	e.bodyData = string(body)
	e.Image.Size = gjson.Get(e.bodyData, "data.size").String()
	e.Image.Filename = gjson.Get(e.bodyData, "data.filename").String()
	e.Image.Url = gjson.Get(e.bodyData, "data.url").String()
	return e
}

// MarkMsgAsRead 标记消息已读
func (h Handle) MarkMsgAsRead(messageId int32) {
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/mark_msg_as_read")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupKick 群组踢人
func (h Handle) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("user_id", string(userId))
	fromData.Add("reject_add_request", fmt.Sprintf("%t", rejectAddRequest))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_kick")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupBan 群组单人禁言
func (h Handle) SetGroupBan(groupId int64, userId int64, duration string) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("user_id", string(userId))
	fromData.Add("duration", duration)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_ban")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupAnonymousBan 群组匿名用户禁言
func (h Handle) SetGroupAnonymousBan(groupId int64, flag string, duration string) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("flag", flag)
	fromData.Add("duration", duration)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_anonymous_ban")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupWholeBan 群组全员禁言
func (h Handle) SetGroupWholeBan(groupId int64, enable bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_whole_ban")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupAdmin 群组设置管理员
func (h Handle) SetGroupAdmin(groupId int64, userId int64, enable bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("userId", string(userId))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_admin")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupAnonymous 群组匿名
func (h Handle) SetGroupAnonymous(groupId int64, enable bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_anonymous")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupName 修改群聊名称
func (h Handle) SetGroupName(groupId int64, groupName string) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("group_name", groupName)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_name")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupLeave 退出群聊
func (h Handle) SetGroupLeave(groupId int64, isDismiss bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("is_dismiss", fmt.Sprintf("%t", isDismiss))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_leave")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupSpecialTitle 设置群组专属头衔
func (h Handle) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration string) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	fromData.Add("user_id", string(userId))
	fromData.Add("special_title", specialTitle)
	fromData.Add("duration", duration)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_special_title")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SendGroupSign 群打卡
func (h Handle) SendGroupSign(groupId int64) {
	fromData := make(url.Values)
	fromData.Add("group_id", string(groupId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_sign")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetFriendAddRequest 处理加好友请求
func (h Handle) SetFriendAddRequest(flag string, approve bool, remark string) {
	fromData := make(url.Values)
	fromData.Add("flag", flag)
	fromData.Add("approve", fmt.Sprintf("%t", approve))
	fromData.Add("remark", remark)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_friend_add_request")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// SetGroupAddRequest 处理加群请求／邀请
func (h Handle) SetGroupAddRequest(flag string, subType string, approve bool, reason string) {
	fromData := make(url.Values)
	fromData.Add("flag", flag)
	fromData.Add("sub_type", subType)
	fromData.Add("approve", fmt.Sprintf("%t", approve))
	fromData.Add("reason", reason)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_add_request")
	http.Post(addr, "application/x-www-form-urlencoded", data)
	//该 API 无响应数据
}

// GetLoginInfo 获取登录号信息
func (h Handle) GetLoginInfo() Event {
	var e Event
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_login_info"))
	if err != nil {
		h.error("获取登录号信息", err)
		return e

	}
	bodyData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		h.error("获取登录号信息", err)
		return e
	}
	e.bodyData = string(bodyData)
	e.Sender.UserID = gjson.Get(e.bodyData, "data.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "nickname").String()
	return e
}
