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

// GetGroupMemberInfo 获取群成员信息
func (h Handle) GetGroupMemberInfo(groupId int64, userId int64, noCache bool) Sender {
	var s Sender
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("no_cache", fmt.Sprintf("%t", noCache))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_member_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return s
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)

	s.GroupID = gjson.Get(bodyData, "data.group_id").String()
	s.UserID = gjson.Get(bodyData, "data.user_id").String()
	s.Nickname = gjson.Get(bodyData, "data.nickname").String()
	s.Card = gjson.Get(bodyData, "data.card").String()
	s.Sex = gjson.Get(bodyData, "data.sex").String()
	s.Age = gjson.Get(bodyData, "data.age").String()
	s.Area = gjson.Get(bodyData, "data.area").String()
	s.JoinTime = gjson.Get(bodyData, "data.join_time").String()
	s.LastSentTime = gjson.Get(bodyData, "data.last_sent_time").String()
	s.Level = gjson.Get(bodyData, "data.level").String()
	s.Role = gjson.Get(bodyData, "data.role").String()
	s.Unfriendly = gjson.Get(bodyData, "data.unfriendly").String()
	s.Title = gjson.Get(bodyData, "data.title").String()
	s.TitleExpireTime = gjson.Get(bodyData, "data.title_expire_time").String()
	s.CardChangeable = gjson.Get(bodyData, "data.card_changeable").String()
	s.ShutUpTimestamp = gjson.Get(bodyData, "data.shut_up_timestamp").String()
	return s
}

// GetGroupMemberList 获取群成员列表
func (h Handle) GetGroupMemberList(groupId int64) []Sender {
	var sender []Sender
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_member_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return sender
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var s Sender
		s.GroupID = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".group_id").String()
		s.UserID = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".user_id").String()
		s.Nickname = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".nickname").String()
		s.Card = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".card").String()
		s.Sex = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".sex").String()
		s.Age = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".age").String()
		s.Area = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".area").String()
		s.JoinTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".join_time").String()
		s.LastSentTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".last_sent_time").String()
		s.Level = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".level").String()
		s.Role = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".role").String()
		s.Unfriendly = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".unfriendly").String()
		s.Title = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".title").String()
		s.TitleExpireTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".title_expire_time").String()
		s.CardChangeable = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".card_changeable").String()
		s.ShutUpTimestamp = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".shut_up_timestamp").String()
		sender = append(sender, s)
	}
	return sender
}

// GetGroupHonorInfo 获取群荣誉信息
func (h Handle) GetGroupHonorInfo(groupId int64, typ string) GroupHonor {
	var gh GroupHonor
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("type", typ)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_honor_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return gh
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	gh.CurrentTalkative.UserId = gjson.Get(bodyData, "data.current_talkative.user_id").String()
	gh.CurrentTalkative.Nickname = gjson.Get(bodyData, "data.current_talkative.nickname").String()
	gh.CurrentTalkative.Avatar = gjson.Get(bodyData, "data.current_talkative.avatar").String()
	gh.CurrentTalkative.DayCount = gjson.Get(bodyData, "data.current_talkative.day_count").String()
	for i := 0; i < int(gjson.Get(bodyData, "data.legend_list.#").Int()); i++ {
		var ho Honor
		ho.UserId = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"user_id").String()
		ho.Nickname = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"nickname").String()
		ho.Avatar = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"avatar").String()
		ho.Description = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"description").String()
		gh.LegendList = append(gh.LegendList, ho)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.performer_list.#").Int()); i++ {
		var ho Honor
		ho.UserId = gjson.Get(bodyData, "data.performer_list."+strconv.Itoa(i)+"user_id").String()
		ho.Nickname = gjson.Get(bodyData, "data.performer_list."+strconv.Itoa(i)+"nickname").String()
		ho.Avatar = gjson.Get(bodyData, "data.performer_list."+strconv.Itoa(i)+"avatar").String()
		ho.Description = gjson.Get(bodyData, "data.performer_list."+strconv.Itoa(i)+"description").String()
		gh.PerformerList = append(gh.PerformerList, ho)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.legend_list.#").Int()); i++ {
		var ho Honor
		ho.UserId = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"user_id").String()
		ho.Nickname = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"nickname").String()
		ho.Avatar = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"avatar").String()
		ho.Description = gjson.Get(bodyData, "data.legend_list."+strconv.Itoa(i)+"description").String()
		gh.LegendList = append(gh.LegendList, ho)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.strong_newbie_list.#").Int()); i++ {
		var ho Honor
		ho.UserId = gjson.Get(bodyData, "data.strong_newbie_list."+strconv.Itoa(i)+"user_id").String()
		ho.Nickname = gjson.Get(bodyData, "data.strong_newbie_list."+strconv.Itoa(i)+"nickname").String()
		ho.Avatar = gjson.Get(bodyData, "data.strong_newbie_list."+strconv.Itoa(i)+"avatar").String()
		ho.Description = gjson.Get(bodyData, "data.strong_newbie_list."+strconv.Itoa(i)+"description").String()
		gh.StrongNewbieList = append(gh.StrongNewbieList, ho)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.emotion_list.#").Int()); i++ {
		var ho Honor
		ho.UserId = gjson.Get(bodyData, "data.emotion_list."+strconv.Itoa(i)+"user_id").String()
		ho.Nickname = gjson.Get(bodyData, "data.emotion_list."+strconv.Itoa(i)+"nickname").String()
		ho.Avatar = gjson.Get(bodyData, "data.emotion_list."+strconv.Itoa(i)+"avatar").String()
		ho.Description = gjson.Get(bodyData, "data.emotion_list."+strconv.Itoa(i)+"description").String()
		gh.EmotionList = append(gh.EmotionList, ho)
	}
	return gh
}

// GetCookies 获取 Cookies
//该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
func (h Handle) GetCookies(domain string) string {
	fromData := make(url.Values)
	fromData.Add("domain", domain)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_cookies")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	cookies := gjson.Get(bodyData, "data.cookies").String()
	return cookies
}

// GetCsrfToken 获取 CSRF Token
//该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
func (h Handle) GetCsrfToken() string {
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_csrf_token"))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	token := gjson.Get(bodyData, "data.token").String()
	return token
}

// GetCredentials 获取 QQ 相关接口凭证
//该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
//即GetCookies和GetCsrfToken两个接口的合并
func (h Handle) GetCredentials(domain string) (string, string) {
	fromData := make(url.Values)
	fromData.Add("domain", domain)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_credentials")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	cookies := gjson.Get(bodyData, "data.cookies").String()
	token := gjson.Get(bodyData, "data.token").String()
	return cookies, token
}

// GetRecord 获取语音
// 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
// 要使用此接口, 通常需要安装 ffmpeg, 请参考 OneBot 实现的相关说明。
func (h Handle) GetRecord(file string, outFormat string) string {
	fromData := make(url.Values)
	fromData.Add("file", file)
	fromData.Add("out_format", outFormat)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_record")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	filepath := gjson.Get(bodyData, "data.file").String()
	return filepath
}

// CanSendImage 检查是否可以发送图片
func (h Handle) CanSendImage() bool {
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/can_send_image"))
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
	yes := gjson.Get(bodyData, "data.yes").Bool()
	return yes

}

// CanSendRecord 检查是否可以发送语音
func (h Handle) CanSendRecord() bool {
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/can_send_record"))
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
	yes := gjson.Get(bodyData, "data.yes").Bool()
	return yes

}

// GetVersionInfo 获取版本信息
func (h Handle) GetVersionInfo() AppInfo {
	var info AppInfo
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_version_info"))
	if err != nil {
		fmt.Println(err)
		return info
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	info.AppName = gjson.Get(bodyData, "data.app_name").String()
	info.AppVersion = gjson.Get(bodyData, "data.app_version").String()
	info.AppFullName = gjson.Get(bodyData, "data.app_full_name").String()
	info.ProtocolVersion = gjson.Get(bodyData, "data.protocol_version").String()
	info.CoolQEdition = gjson.Get(bodyData, "data.coolq_edition").String()
	info.CoolQDirectory = gjson.Get(bodyData, "data.coolq_directory").String()
	info.GoCqHttp = gjson.Get(bodyData, "data.go-cqhttp").String()
	info.PluginVersion = gjson.Get(bodyData, "data.plugin_version").String()
	info.PluginBuildNumber = gjson.Get(bodyData, "data.plugin_build_number").String()
	info.PluginBuildConfiguration = gjson.Get(bodyData, "data.plugin_build_configuration").String()
	info.RuntimeVersion = gjson.Get(bodyData, "data.runtime_version").String()
	info.RuntimeOs = gjson.Get(bodyData, "data.runtime_os").String()
	info.Version = gjson.Get(bodyData, "data.version").String()
	info.Protocol = gjson.Get(bodyData, "data.protocol").String()
	return info
}

// SetRestart 重启 go-cqhttp
func (h Handle) SetRestart(delay int) {
	fromData := make(url.Values)
	fromData.Add("delay", strconv.Itoa(delay))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_restart")
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
	// 该 API 无响应数据
}

// CleanCache 清理缓存
func (h Handle) CleanCache() {
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/clean_cache"))
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
	// 该 API 无响应数据
}

// SetGroupPortrait 设置群头像
// 目前这个API在登录一段时间后因cookie失效而失效, 请考虑后使用
func (h Handle) SetGroupPortrait(groupId int64, file string, cache int) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file", file)
	fromData.Add("cache", strconv.Itoa(cache))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_group_portrait")
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

// GetWordSlices 获取中文分词 ( 隐藏 API )
// 警告:隐藏 API 是不建议一般用户使用的, 它们只应该在 OneBot 实现内部或由 SDK 和框架使用, 因为不正确的使用可能造成程序运行不正常。
//func (h Handle) GetWordSlices(content string) []string {
//	var slices []string
//	fromData := make(url.Values)
//	fromData.Add("content", content)
//	data := strings.NewReader(fromData.Encode())
//	addr := fmt.Sprintf("http://" + h.Host + "/.get_word_slices")
//	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
//	if err != nil {
//		fmt.Println(err)
//		return slices
//
//	}
//	bodyData := h.noData(request.Body)
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(request.Body)
//	for i := 0; i < int(gjson.Get(bodyData, "data.slices.#").Int()); i++ {
//		s := gjson.Get(bodyData, "data.slices."+strconv.Itoa(i)).String()
//		slices = append(slices, s)
//	}
//	return slices
//}

// OcrImage 图片 OCR
// 目前图片OCR接口仅支持接受的图片
func (h Handle) OcrImage(image string) ImageOCR {
	var ocr ImageOCR
	fromData := make(url.Values)
	fromData.Add("image", image)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/ocr_image")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return ocr

	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	ocr.Language = gjson.Get(bodyData, "data.language").String()
	for i := 0; i < int(gjson.Get(bodyData, "data.texts.#").Int()); i++ {
		var texts TextDetection
		texts.Test = gjson.Get(bodyData, "data.texts."+strconv.Itoa(i)+".text").String()
		texts.Confidence = gjson.Get(bodyData, "data.texts."+strconv.Itoa(i)+".confidence").String()
		texts.Coordinates = gjson.Get(bodyData, "data.texts."+strconv.Itoa(i)+".coordinates").String()
		ocr.Texts = append(ocr.Texts, texts)
	}
	return ocr
}

// GetGroupSystemMsg 获取群系统消息
// 如果列表不存在任何消息, 将返回 null
func (h Handle) GetGroupSystemMsg() ([]InvitedRequest, []JoinRequest) {
	var invited []InvitedRequest
	var join []JoinRequest
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_group_system_msg"))
	if err != nil {
		fmt.Println(err)
		return invited, join
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.invited_requests.#").Int()); i++ {
		var inv InvitedRequest
		inv.RequestId = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".request_id").String()
		inv.InvitorUin = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".invitor_uin").String()
		inv.InvitorNick = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".invitor_nick").String()
		inv.GroupId = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".group_id").String()
		inv.GroupName = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".group_name").String()
		inv.Checked = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".checked").String()
		inv.Actor = gjson.Get(bodyData, "data.invited_requests."+strconv.Itoa(i)+".actor").String()
		invited = append(invited, inv)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.join_requests.#").Int()); i++ {
		var jo JoinRequest
		jo.RequestId = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".request_id").String()
		jo.RequesterUin = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".requester_uin").String()
		jo.RequesterNick = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".requester_nick").String()
		jo.Message = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".message").String()
		jo.GroupId = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".group_id").String()
		jo.GroupName = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".group_name").String()
		jo.Checked = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".checked").String()
		jo.Actor = gjson.Get(bodyData, "data.join_requests."+strconv.Itoa(i)+".actor").String()
		join = append(join, jo)
	}
	return invited, join
}

// UploadPrivateFile 上传私聊文件
func (h Handle) UploadPrivateFile(userId int64, file string, name string) {
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("file", file)
	fromData.Add("name", name)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/upload_private_file")
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

// UploadGroupFile 上传群文件
func (h Handle) UploadGroupFile(groupId int64, file string, name string, folder string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file", file)
	fromData.Add("name", name)
	fromData.Add("folder", folder)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/upload_group_file")
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

// GetGroupFileSystemInfo 获取群文件系统信息
func (h Handle) GetGroupFileSystemInfo(groupId int64) GroupFileSystemInfo {
	var g GroupFileSystemInfo
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_file_system_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return g
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	g.FileCount = gjson.Get(bodyData, "data.file_count").String()
	g.LimitCount = gjson.Get(bodyData, "data.limit_count").String()
	g.UsedSpace = gjson.Get(bodyData, "data.used_space").String()
	g.TotalSpace = gjson.Get(bodyData, "data.total_space").String()
	return g

}

// GetGroupRootFiles 获取群根目录文件列表
func (h Handle) GetGroupRootFiles(groupId int64) ([]GroupFile, []Folder) {
	var files []GroupFile
	var folders []Folder
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_root_files")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return files, folders
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.files.#").Int()); i++ {
		var f GroupFile
		f.GroupId = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".group_id").String()
		f.FileId = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".file_id").String()
		f.FileName = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".file_name").String()
		f.BusId = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".busid").String()
		f.FileSize = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".file_size").String()
		f.UploadTime = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".upload_time").String()
		f.DeadTime = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".dead_time").String()
		f.ModifyTime = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".modify_time").String()
		f.DownloadTimes = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".download_times").String()
		f.Uploader = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".uploader").String()
		f.UploaderName = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".uploader_name").String()
		files = append(files, f)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.folders.#").Int()); i++ {
		var f Folder
		f.GroupId = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".group_id").String()
		f.FolderId = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".folder_id").String()
		f.FolderName = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".folder_name").String()
		f.CreateTime = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".create_time").String()
		f.Creator = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".creator").String()
		f.CreatorName = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".creator_name").String()
		f.TotalFileCount = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".total_file_count").String()
		folders = append(folders, f)
	}
	return files, folders
}

// GetGroupFilesByFolder 获取群子目录文件列表
func (h Handle) GetGroupFilesByFolder(groupId int64, folderId string) ([]GroupFile, []Folder) {
	var files []GroupFile
	var folders []Folder
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("folder_id", folderId)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_files_by_folder")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return files, folders
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.files.#").Int()); i++ {
		var f GroupFile
		f.GroupId = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".group_id").String()
		f.FileId = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".file_id").String()
		f.FileName = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".file_name").String()
		f.BusId = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".busid").String()
		f.FileSize = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".file_size").String()
		f.UploadTime = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".upload_time").String()
		f.DeadTime = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".dead_time").String()
		f.ModifyTime = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".modify_time").String()
		f.DownloadTimes = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".download_times").String()
		f.Uploader = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".uploader").String()
		f.UploaderName = gjson.Get(bodyData, "data.files."+strconv.Itoa(i)+".uploader_name").String()
		files = append(files, f)
	}
	for i := 0; i < int(gjson.Get(bodyData, "data.folders.#").Int()); i++ {
		var f Folder
		f.GroupId = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".group_id").String()
		f.FolderId = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".folder_id").String()
		f.FolderName = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".folder_name").String()
		f.CreateTime = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".create_time").String()
		f.Creator = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".creator").String()
		f.CreatorName = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".creator_name").String()
		f.TotalFileCount = gjson.Get(bodyData, "data.folders."+strconv.Itoa(i)+".total_file_count").String()
		folders = append(folders, f)
	}
	return files, folders
}

// CreateGroupFileFolder 创建群文件文件夹
func (h Handle) CreateGroupFileFolder(groupId int64, name string, parentId string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("name", name)
	fromData.Add("parent_id", parentId)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/create_group_file_folder")
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

// DeleteGroupFolder 删除群文件夹
func (h Handle) DeleteGroupFolder(groupId int64, folderId string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("folder_id", folderId)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/delete_group_folder")
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

// DeleteGroupFile 删除群文件
func (h Handle) DeleteGroupFile(groupId int64, fileId string, busId int32) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file_id", fileId)
	fromData.Add("busid", fmt.Sprint(busId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/delete_group_file")
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

// GetGroupFileUrl 获取群文件资源链接
func (h Handle) GetGroupFileUrl(groupId int64, fileId string, busId int32) string {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file_id", fileId)
	fromData.Add("busid", fmt.Sprint(busId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_file_url")
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
	return gjson.Get(bodyData, "data.url").String()
}

// GetStatus 获取状态
func (h Handle) GetStatus() Status {
	var status Status
	request, err := http.Get(fmt.Sprintf("http://" + h.Host + "/get_status"))
	if err != nil {
		fmt.Println(err)
		return status
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	status.AppInitialized = gjson.Get(bodyData, "data.app_initialized").String()
	status.AppEnabled = gjson.Get(bodyData, "data.app_enabled").String()
	status.PluginsGood = gjson.Get(bodyData, "data.plugins_good").String()
	status.AppGood = gjson.Get(bodyData, "data.app_good").String()
	status.Online = gjson.Get(bodyData, "data.online").String()
	status.Good = gjson.Get(bodyData, "data.good").String()
	status.Stat.PacketReceived = gjson.Get(bodyData, "data.stat.PacketReceived").String()
	status.Stat.PacketSent = gjson.Get(bodyData, "data.stat.PacketSent").String()
	status.Stat.PacketLost = gjson.Get(bodyData, "data.stat.PacketLost").String()
	status.Stat.MessageReceived = gjson.Get(bodyData, "data.stat.MessageReceived").String()
	status.Stat.MessageSent = gjson.Get(bodyData, "data.stat.MessageSent").String()
	status.Stat.DisconnectTimes = gjson.Get(bodyData, "data.stat.DisconnectTimes").String()
	status.Stat.LostTimes = gjson.Get(bodyData, "data.stat.LostTimes").String()
	status.Stat.LastMessageTime = gjson.Get(bodyData, "data.stat.LastMessageTime").String()
	return status
}

// GetGroupAtAllRemain 获取群@全体成员剩余次数
func (h Handle) GetGroupAtAllRemain(groupId int64) GroupAtAllRemain {
	var g GroupAtAllRemain
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_at_all_remain")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return g
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	g.CanAtAll = gjson.Get(bodyData, "data.can_at_all").String()
	g.RemainAtAllCountForGroup = gjson.Get(bodyData, "data.remain_at_all_count_for_group").String()
	g.RemainAtAllCountForUin = gjson.Get(bodyData, "data.remain_at_all_count_for_uin").String()
	return g
}

// HandleQuickOperation 对事件执行快速操作 ( 隐藏 API )
//func (h Handle)HandleQuickOperation (){
//
//}

// SendGroupNotice 发送群公告
func (h Handle) SendGroupNotice(groupId int64, content string, image string) {
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("content", content)
	fromData.Add("image", image)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/_send_group_notice")
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

// GetGroupNotice 获取群公告
func (h Handle) GetGroupNotice(groupId int64) []GroupNotice {
	var gn []GroupNotice
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/_get_group_notice")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return gn
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var g GroupNotice
		g.SenderId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".sender_id").String()
		g.PublishTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".message.text").String()
		g.Message.Text = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".message.text").String()
		g.Message.Images.Height = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".message.images.height").String()
		g.Message.Images.Width = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".message.images.width").String()
		g.Message.Images.Id = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".message.images.id").String()
		gn = append(gn, g)
	}
	return gn
}

// ReloadEventFilter 重载事件过滤器
func (h Handle) ReloadEventFilter(file string) {
	fromData := make(url.Values)
	fromData.Add("file", file)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/reload_event_filter")
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

// DownloadFile 下载文件到缓存目录
func (h Handle) DownloadFile(fileUrl string, threadCount int32, headers string) string {
	fromData := make(url.Values)
	fromData.Add("url", fileUrl)
	fromData.Add("thread_count", fmt.Sprint(threadCount))
	fromData.Add("headers", headers)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/download_file")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	return gjson.Get(bodyData, "data.file").String()
}

// GetOnlineClients 获取当前账号在线客户端列表
func (h Handle) GetOnlineClients(noCache bool) []Clients {
	var clients []Clients
	fromData := make(url.Values)
	fromData.Add("no_cache", fmt.Sprint(noCache))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_online_clients")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return clients
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var client Clients
		client.AppId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+"app_id").String()
		client.AppId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+"app_id").String()
		client.AppId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+"app_id").String()
		clients = append(clients, client)
	}
	return clients
}

// GetGroupMsgHistory 获取群消息历史记录
func (h Handle) GetGroupMsgHistory(messageSeq int64, groupId int64) string {
	fromData := make(url.Values)
	fromData.Add("message_seq", strconv.FormatInt(messageSeq, 10))
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_group_msg_history")
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
	return bodyData
}

// SetEssenceMsg 设置精华消息
func (h Handle) SetEssenceMsg(messageId int32) {
	fromData := make(url.Values)
	fromData.Add("message_id", fmt.Sprint(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/set_essence_msg")
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

// DeleteEssenceMsg 移除精华消息
func (h Handle) DeleteEssenceMsg(messageId int32) {
	fromData := make(url.Values)
	fromData.Add("message_id", fmt.Sprint(messageId))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/delete_essence_msg")
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

// GetEssenceMsgList 获取精华消息列表
func (h Handle) GetEssenceMsgList(groupId int64) []EssenceMsg {
	var EssenceMsgList []EssenceMsg
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/get_essence_msg_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return EssenceMsgList
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var essenceMsg EssenceMsg
		essenceMsg.SenderId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".sender_id").String()
		essenceMsg.SenderNick = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".sender_nick").String()
		essenceMsg.SenderTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".sender_time").String()
		essenceMsg.OperatorId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".operator_id").String()
		essenceMsg.OperatorNick = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".operator_nick").String()
		essenceMsg.OperatorTime = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".operator_time").String()
		essenceMsg.MessageId = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".message_id").String()
		EssenceMsgList = append(EssenceMsgList, essenceMsg)
	}
	return EssenceMsgList
}

// CheckUrlSafely 检查链接安全性
func (h Handle) CheckUrlSafely(u string) string {
	fromData := make(url.Values)
	fromData.Add("url", u)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/check_url_safely")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	return gjson.Get(bodyData, "data.level").String()
}
func (h Handle) GetModelShow(model string) []Variant {
	var variants []Variant
	fromData := make(url.Values)
	fromData.Add("model", model)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/_get_model_show")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	bodyData := h.noData(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	for i := 0; i < int(gjson.Get(bodyData, "data.#").Int()); i++ {
		var variant Variant
		variant.ModelShow = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".model_show").String()
		variant.NeedPay = gjson.Get(bodyData, "data."+strconv.Itoa(i)+".need_pay").String()
		variants = append(variants, variant)
	}
	return variants

}

// SetModelShow 设置在线机型
func (h Handle) SetModelShow(model string, modelShow string) {
	fromData := make(url.Values)
	fromData.Add("model", model)
	fromData.Add("model_show", modelShow)
	data := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf("http://" + h.Host + "/_set_model_show")
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
