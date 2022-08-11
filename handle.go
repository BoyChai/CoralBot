package CoralBot

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Handle struct {
	Host      string
	Agreement string
}

// SendPrivateMsg 发送私聊消息
func (h Handle) SendPrivateMsg(userId int64, groupId int64, message string, autoEscape bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("message", message)
	fromData.Add("auto_escape", fmt.Sprint(autoEscape))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_private_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SendGroupMsg 发送群聊消息
func (h Handle) SendGroupMsg(groupId int64, message string, autoEscape string) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_group_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SendGroupForwardMsg 发送合并转发 ( 群 )
//func (h Handle) SendGroupForwardMsg(groupId int64, message string) {
//	var e Event
//	fromData := make(url.Values)
//	fromData.Add("group_id", groupId)
//	fromData.Add("message", message)
//	data := strings.NewReader(fromData.Encode())
//	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_group_forward_msg")
//	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
//	if err != nil {
//		fmt.Println(err)
//		return e
//	}
//	h.noData(request.Body)
//	defer func(Body io.ReadCloser) {
//		err = Body.Close()
//		fmt.Println(err)
//	}(request.Body)
//}

// SendMsg 发送消息
func (h Handle) SendMsg(userId int64, groupId int64, message string, autoEscape string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("message", message)
	fromData.Add("auto_escape", autoEscape)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DeleteMsg 撤回消息
func (h Handle) DeleteMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetMsg 获取消息
func (h Handle) GetMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetForwardMsg 获取合并转发内容
func (h Handle) GetForwardMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_forward_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetImage 获取图片信息
func (h Handle) GetImage(file string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("file", file)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_image")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// MarkMsgAsRead 标记消息已读
func (h Handle) MarkMsgAsRead(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/mark_msg_as_read")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupKick 群组踢人
func (h Handle) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("reject_add_request", fmt.Sprintf("%t", rejectAddRequest))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_kick")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupBan 群组单人禁言
func (h Handle) SetGroupBan(groupId int64, userId int64, duration string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("duration", duration)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_ban")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupAnonymousBan 群组匿名用户禁言
func (h Handle) SetGroupAnonymousBan(groupId int64, flag string, duration string) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("flag", flag)
	fromData.Add("duration", duration)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_anonymous_ban")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupWholeBan 群组全员禁言
func (h Handle) SetGroupWholeBan(groupId int64, enable bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_whole_ban")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupAdmin 群组设置管理员
func (h Handle) SetGroupAdmin(groupId int64, userId int64, enable bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("userId", strconv.FormatInt(userId, 10))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_admin")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupAnonymous 群组匿名
func (h Handle) SetGroupAnonymous(groupId int64, enable bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("enable", fmt.Sprintf("%t", enable))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_anonymous")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupName 修改群聊名称
func (h Handle) SetGroupName(groupId int64, groupName string) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("group_name", groupName)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_name")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupLeave 退出群聊
func (h Handle) SetGroupLeave(groupId int64, isDismiss bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("is_dismiss", fmt.Sprintf("%t", isDismiss))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_leave")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupSpecialTitle 设置群组专属头衔
func (h Handle) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration string) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("special_title", specialTitle)
	fromData.Add("duration", duration)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_special_title")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SendGroupSign 群打卡
func (h Handle) SendGroupSign(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_group_sign")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetFriendAddRequest 处理加好友请求
func (h Handle) SetFriendAddRequest(flag string, approve bool, remark string) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("flag", flag)
	fromData.Add("approve", fmt.Sprintf("%t", approve))
	fromData.Add("remark", remark)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_friend_add_request")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupAddRequest 处理加群请求／邀请
func (h Handle) SetGroupAddRequest(flag string, subType string, approve bool, reason string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("flag", flag)
	fromData.Add("sub_type", subType)
	fromData.Add("approve", fmt.Sprintf("%t", approve))
	fromData.Add("reason", reason)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_add_request")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetLoginInfo 获取登录号信息
func (h Handle) GetLoginInfo() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_login_info"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// QiDianGetAccountInfo 获取企点账号信息
// 该API只有企点协议可用
func (h Handle) QiDianGetAccountInfo() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/qidian_get_account_info"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetQQProfile 设置登录号资料
func (h Handle) SetQQProfile(p Profile) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("nickname", p.Nickname)
	fromData.Add("company", p.Company)
	fromData.Add("email", p.Email)
	fromData.Add("college", p.College)
	fromData.Add("personal_note", p.PersonalNote)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_qq_profile")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetStrangerInfo 获取陌生人信息
func (h Handle) GetStrangerInfo(userId int64, noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("no_cache", fmt.Sprintf("%t", noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_stranger_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetFriendList 获取好友列表
func (h Handle) GetFriendList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_friend_list"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetUnidirectionalFriendList 获取单向好友列表
func (h Handle) GetUnidirectionalFriendList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_unidirectional_friend_list"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DeleteFriend 删除好友
func (h Handle) DeleteFriend(friendId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("friend_id", strconv.FormatInt(friendId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_friend")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupInfo 获取群信息
func (h Handle) GetGroupInfo(groupId int64, noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("no_cache", fmt.Sprintf("%t", noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupList 获取群列表
func (h Handle) GetGroupList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_unidirectional_friend_list"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupMemberInfo 获取群成员信息
func (h Handle) GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("no_cache", fmt.Sprintf("%t", noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_member_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupMemberList 获取群成员列表
func (h Handle) GetGroupMemberList(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_member_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupHonorInfo 获取群荣誉信息
func (h Handle) GetGroupHonorInfo(groupId int64, typ string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("type", typ)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_honor_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetCookies 获取 Cookies
//该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
func (h Handle) GetCookies(domain string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("domain", domain)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_cookies")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetCsrfToken 获取 CSRF Token
//该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
func (h Handle) GetCsrfToken() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_csrf_token"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetCredentials 获取 QQ 相关接口凭证
//该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
//即GetCookies和GetCsrfToken两个接口的合并
func (h Handle) GetCredentials(domain string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("domain", domain)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_credentials")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetRecord 获取语音
// 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
// 要使用此接口, 通常需要安装 ffmpeg, 请参考 OneBot 实现的相关说明。
func (h Handle) GetRecord(file string, outFormat string) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("file", file)
	fromData.Add("out_format", outFormat)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_record")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// CanSendImage 检查是否可以发送图片
func (h Handle) CanSendImage() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/can_send_image"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// CanSendRecord 检查是否可以发送语音
func (h Handle) CanSendRecord() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/can_send_record"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetVersionInfo 获取版本信息
func (h Handle) GetVersionInfo() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_version_info"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetRestart 重启 go-cqhttp
func (h Handle) SetRestart(delay int) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("delay", strconv.Itoa(delay))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_restart")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// CleanCache 清理缓存
func (h Handle) CleanCache() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/clean_cache"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetGroupPortrait 设置群头像
// 目前这个API在登录一段时间后因cookie失效而失效, 请考虑后使用
func (h Handle) SetGroupPortrait(groupId int64, file string, cache int) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file", file)
	fromData.Add("cache", strconv.Itoa(cache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_group_portrait")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetWordSlices 获取中文分词 ( 隐藏 API )
// 警告:隐藏 API 是不建议一般用户使用的, 它们只应该在 OneBot 实现内部或由 SDK 和框架使用, 因为不正确的使用可能造成程序运行不正常。
//func (h Handle) GetWordSlices(content string) []string {
//	var slices []string
//	fromData := make(url.Values)
//	fromData.Add("content", content)
//	data := strings.NewReader(fromData.Encode())
//	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/.get_word_slices")
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
func (h Handle) OcrImage(image string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("image", image)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/ocr_image")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupSystemMsg 获取群系统消息
// 如果列表不存在任何消息, 将返回 null
func (h Handle) GetGroupSystemMsg() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_system_msg"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// UploadPrivateFile 上传私聊文件
func (h Handle) UploadPrivateFile(userId int64, file string, name string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(userId, 10))
	fromData.Add("file", file)
	fromData.Add("name", name)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/upload_private_file")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// UploadGroupFile 上传群文件
func (h Handle) UploadGroupFile(groupId int64, file string, name string, folder string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file", file)
	fromData.Add("name", name)
	fromData.Add("folder", folder)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/upload_group_file")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupFileSystemInfo 获取群文件系统信息
func (h Handle) GetGroupFileSystemInfo(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_file_system_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupRootFiles 获取群根目录文件列表
func (h Handle) GetGroupRootFiles(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_root_files")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupFilesByFolder 获取群子目录文件列表
func (h Handle) GetGroupFilesByFolder(groupId int64, folderId string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("folder_id", folderId)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_files_by_folder")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// CreateGroupFileFolder 创建群文件文件夹
func (h Handle) CreateGroupFileFolder(groupId int64, name string, parentId string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("name", name)
	fromData.Add("parent_id", parentId)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/create_group_file_folder")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DeleteGroupFolder 删除群文件夹
func (h Handle) DeleteGroupFolder(groupId int64, folderId string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("folder_id", folderId)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_group_folder")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DeleteGroupFile 删除群文件
func (h Handle) DeleteGroupFile(groupId int64, fileId string, busId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file_id", fileId)
	fromData.Add("busid", fmt.Sprint(busId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_group_file")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupFileUrl 获取群文件资源链接
func (h Handle) GetGroupFileUrl(groupId int64, fileId string, busId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("file_id", fileId)
	fromData.Add("busid", fmt.Sprint(busId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_file_url")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetStatus 获取状态
func (h Handle) GetStatus() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_status"))
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupAtAllRemain 获取群@全体成员剩余次数
func (h Handle) GetGroupAtAllRemain(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_at_all_remain")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// HandleQuickOperation 对事件执行快速操作 ( 隐藏 API )
//func (h Handle)HandleQuickOperation (){
//
//}

// SendGroupNotice 发送群公告
func (h Handle) SendGroupNotice(groupId int64, content string, image string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	fromData.Add("content", content)
	fromData.Add("image", image)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/_send_group_notice")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupNotice 获取群公告
func (h Handle) GetGroupNotice(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/_get_group_notice")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// ReloadEventFilter 重载事件过滤器
func (h Handle) ReloadEventFilter(file string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("file", file)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/reload_event_filter")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DownloadFile 下载文件到缓存目录
func (h Handle) DownloadFile(fileUrl string, threadCount int32, headers string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("url", fileUrl)
	fromData.Add("thread_count", fmt.Sprint(threadCount))
	fromData.Add("headers", headers)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/download_file")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetOnlineClients 获取当前账号在线客户端列表
func (h Handle) GetOnlineClients(noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("no_cache", fmt.Sprint(noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_online_clients")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGroupMsgHistory 获取群消息历史记录
func (h Handle) GetGroupMsgHistory(messageSeq int64, groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_seq", strconv.FormatInt(messageSeq, 10))
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_msg_history")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetEssenceMsg 设置精华消息
func (h Handle) SetEssenceMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", fmt.Sprint(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_essence_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DeleteEssenceMsg 移除精华消息
func (h Handle) DeleteEssenceMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", fmt.Sprint(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_essence_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetEssenceMsgList 获取精华消息列表
func (h Handle) GetEssenceMsgList(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_essence_msg_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// CheckUrlSafely 检查链接安全性
func (h Handle) CheckUrlSafely(u string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("url", u)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/check_url_safely")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (h Handle) GetModelShow(model string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("model", model)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/_get_model_show")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// SetModelShow 设置在线机型
func (h Handle) SetModelShow(model string, modelShow string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("model", model)
	fromData.Add("model_show", modelShow)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/_set_model_show")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		fmt.Println(err)
	}(request.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
