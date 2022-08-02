package CoralBot

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type Handle struct {
	Host string
}

func (h Handle) noData(body io.ReadCloser) string {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	status := gjson.Get(string(data), "status").String()
	if status != "ok" {
		fmt.Println("执行失败")

	}
	return string(data)
}

// SendPrivateMsg 发送私聊消息
func (h Handle) SendPrivateMsg(userId string, groupId string, message string, autoEscape string) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("user_id", userId)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_private_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.MessageId = gjson.Get(bodyData, "data.message_id").String()
	return e
}

// SendGroupMsg 发送群聊消息
func (h Handle) SendGroupMsg(groupId string, message string, autoEscape string) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.MessageId = gjson.Get(bodyData, "data.message_id").String()
	return e
}

// SendGroupForwardMsg 发送合并转发 ( 群 )
//func (h Handle) SendGroupForwardMsg(groupId string, message string) {
//	var e Event
//	fromData := make(url.Values)
//	fromData.Add("group_id", groupId)
//	fromData.Add("message", message)
//	data := strings.NewReader(fromData.Encode())
//	addr := fmt.Sprintf("http://" + h.Host + "/send_group_forward_msg")
//	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
//	if err != nil {
//		fmt.Println(err)
//		return e
//	}
//	h.noData(request.Body)
//	defer request.Body.Close()
//}

// SendMsg 发送消息
func (h Handle) SendMsg(userId string, groupId string, message string, autoEscape string) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("user_id", userId)
	fromData.Add("group_id", groupId)
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.MessageId = gjson.Get(bodyData, "data.message_id").String()
	return e
}

// DeleteMsg 撤回消息
func (h Handle) DeleteMsg(messageId int32) {
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/delete_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
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
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.bodyData = bodyData
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
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.bodyData = bodyData
	for i := 0; i < int(gjson.Get(e.bodyData, "data.messages.#").Int()); i++ {
		var messages ForwardMessage
		messages.Content = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+strconv.Itoa(i)+".content")).String()
		messages.Sender.Nickname = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+strconv.Itoa(i)+".sender.nickname")).String()
		messages.Sender.UserID = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+strconv.Itoa(i)+".sender.user_id")).String()
		messages.Time = gjson.Get(e.bodyData, fmt.Sprintf("data.messages."+strconv.Itoa(i)+".time")).String()
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
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.bodyData = bodyData
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
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupKick 群组踢人
func (h Handle) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("reject_add_request", fmt.Sprintf("%t", rejectAddRequest))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_kick")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupBan 群组单人禁言
func (h Handle) SetGroupBan(groupId int64, userId int64, duration string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("duration", duration)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_ban")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupAnonymousBan 群组匿名用户禁言
func (h Handle) SetGroupAnonymousBan(groupId int64, flag string, duration string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("flag", flag)
	fromData.Add("duration", duration)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_anonymous_ban")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupWholeBan 群组全员禁言
func (h Handle) SetGroupWholeBan(groupId int64, enable bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_whole_ban")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupAdmin 群组设置管理员
func (h Handle) SetGroupAdmin(groupId int64, userId int64, enable bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("userId", strconv.FormatInt(userId, 10))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_admin")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupAnonymous 群组匿名
func (h Handle) SetGroupAnonymous(groupId int64, enable bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_anonymous")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupName 修改群聊名称
func (h Handle) SetGroupName(groupId int64, groupName string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("group_name", groupName)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_name")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupLeave 退出群聊
func (h Handle) SetGroupLeave(groupId int64, isDismiss bool) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("is_dismiss", fmt.Sprintf("%t", isDismiss))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_leave")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SetGroupSpecialTitle 设置群组专属头衔
func (h Handle) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("special_title", specialTitle)
	fromData.Add("duration", duration)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_special_title")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// SendGroupSign 群打卡
func (h Handle) SendGroupSign(groupId int64) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/send_group_sign")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
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
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
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
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// GetLoginInfo 获取登录号信息
func (h Handle) GetLoginInfo() Event {
	var e Event
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_login_info"))
	if err != nil {
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.bodyData = bodyData
	e.Sender.UserID = gjson.Get(e.bodyData, "data.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "nickname").String()
	return e
}

// QiDianGetAccountInfo 获取企点账号信息
// 该API只有企点协议可用
func (h Handle) QiDianGetAccountInfo() Event {
	var e Event
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/qidian_get_account_info"))
	if err != nil {
		fmt.Println(err)
		return e
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.bodyData = bodyData
	e.QiDian.MasterId = gjson.Get(e.bodyData, "master_id").String()
	e.QiDian.ExtName = gjson.Get(e.bodyData, "ext_name").String()
	e.QiDian.CreateTime = gjson.Get(e.bodyData, "create_time").String()
	return e
}

// SetQQProfile 设置登录号资料
func (h Handle) SetQQProfile(p Profile) {
	fromData := make(url.Values)
	fromData.Add("nickname", p.Nickname)
	fromData.Add("company", p.Company)
	fromData.Add("email", p.Email)
	fromData.Add("college", p.College)
	fromData.Add("personal_note", p.PersonalNote)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_qq_profile")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
}

// GetStrangerInfo 获取陌生人信息
func (h Handle) GetStrangerInfo(userId int64, noCache bool) Event {
	var e Event
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("no_cache", fmt.Sprintf("%t", noCache))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_stranger_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	e.bodyData = bodyData
	e.Sender.UserID = gjson.Get(e.bodyData, "data.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "data.nickname").String()
	e.Sender.Sex = gjson.Get(e.bodyData, "data.sex").String()
	e.Sender.Age = gjson.Get(e.bodyData, "data.age").String()
	e.Sender.QID = gjson.Get(e.bodyData, "data.qid").String()
	e.Sender.Level = gjson.Get(e.bodyData, "data.level").String()
	e.Sender.LoginDays = gjson.Get(e.bodyData, "data.login_days").String()
	return e
}

// GetFriendList 获取好友列表
func (h Handle) GetFriendList() []Sender {
	var s []Sender
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_friend_list"))
	if err != nil {
		fmt.Println(err)
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var sender Sender
		sender.UserID = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".user_id").String()
		sender.Nickname = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".nickname").String()
		sender.Remark = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".remark").String()
		s = append(s, sender)
	}
	return s
}

// GetUnidirectionalFriendList 获取单向好友列表
func (h Handle) GetUnidirectionalFriendList() []Sender {
	var s []Sender
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_unidirectional_friend_list"))
	if err != nil {
		fmt.Println(err)
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var sender Sender
		sender.UserID = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".user_id").String()
		sender.Nickname = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".nickname").String()
		sender.Source = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".source").String()
		s = append(s, sender)
	}
	return s
}

// DeleteFriend 删除好友
func (h Handle) DeleteFriend(friendId int64) {
	fromData := make(url.Values)
	fromData.Add("friend_id", strconv.FormatInt(friendId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/delete_friend")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	//该 API 无响应数据
}

// GetGroupInfo 获取群信息
func (h Handle) GetGroupInfo(groupId int64, noCache bool) GroupInfo {
	var g GroupInfo
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("no_cache", fmt.Sprintf("%t", noCache))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	g.GroupId = gjson.Get(bodyData, "data.group_id").String()
	g.GroupName = gjson.Get(bodyData, "data.group_name").String()
	g.GroupMemo = gjson.Get(bodyData, "data.group_memo").String()
	g.GroupCreateTime = gjson.Get(bodyData, "data.group_create_time").String()
	g.GroupLevel = gjson.Get(bodyData, "data.group_level").String()
	g.MemberCount = gjson.Get(bodyData, "data.member_count").String()
	g.MaxMemberCount = gjson.Get(bodyData, "data.max_member_count").String()
	return g
}

// GetGroupList 获取群列表
func (h Handle) GetGroupList() []GroupInfo {
	var groupList []GroupInfo
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_unidirectional_friend_list"))
	if err != nil {
		fmt.Println(err)
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var g GroupInfo
		g.GroupId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".group_id").String()
		g.GroupName = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".group_name").String()
		g.GroupMemo = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".group_memo").String()
		g.GroupCreateTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".group_create_time").String()
		g.GroupLevel = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".group_level").String()
		g.MemberCount = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".member_count").String()
		g.MaxMemberCount = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".max_member_count").String()
		groupList = append(groupList, g)
	}
	return groupList
}
