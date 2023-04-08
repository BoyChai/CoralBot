package action

import (
	"encoding/json"
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type QQHandle struct {
	Host      string
	Agreement string
}
type QQMsg struct {
	UserId     int64
	GroupId    int64
	Message    string
	AutoEscape bool
}

// Reply 快速回复
func (h QQHandle) Reply(e bot.QQEvent, m QQMsg) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch e.MessageType {
	case "group":
		m.GroupId = e.GroupID
		data, err = h.SendMsg(m)
	case "private":
		m.UserId = e.UserID
		data, err = h.SendMsg(m)
	case "guild":
		data, err = h.SendGuildChannelMsg(e.GuildID, e.ChannelID, m.Message)
	}
	return data, err

}

// SendPrivateMsg 发送私聊消息
func (h QQHandle) SendPrivateMsg(m QQMsg) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(m.UserId, 10))
	fromData.Add("group_id", strconv.FormatInt(m.GroupId, 10))
	fromData.Add("message", m.Message)
	fromData.Add("auto_escape", fmt.Sprint(m.AutoEscape))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_private_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}

	body, err := io.ReadAll(request.Body)
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
// func (h QQHandle) SendGroupMsg(groupId int64, message string, autoEscape string) (map[string]interface{}, error) {
func (h QQHandle) SendGroupMsg(m QQMsg) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(m.GroupId, 10))
	fromData.Add("message", m.Message)
	fromData.Add("auto_escape", fmt.Sprint(m.AutoEscape))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_group_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
	if err != nil {
		return data, err
	}
	body, err := io.ReadAll(request.Body)
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

//SendGroupForwardMsg 发送合并转发(群)
//func (h QQHandle) SendGroupForwardMsg(groupId int64, message string) (map[string]interface{}, error) {
//	var data map[string]interface{}
//	fromData := make(url.Values)
//	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
//	fromData.Add("message", message)
//	readerData := strings.NewReader(fromData.Encode())
//	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_group_forward_msg")
//	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
//	if err != nil {
//		return data, err
//	}
//	body, err := io.ReadAll(request.Body)
//	if err != nil {
//		return data, err
//	}
//	defer func(Body io.ReadCloser) {
//		err = Body.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(request.Body)
//	err = json.Unmarshal(body, &data)
//	if err != nil {
//		return data, err
//	}
//	return data, nil
//}

// SendMsg 发送消息
// func (h QQHandle) SendMsg(userId int64, groupId int64, message string, autoEscape string) (map[string]interface{}, error) {
func (h QQHandle) SendMsg(m QQMsg) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("user_id", strconv.FormatInt(m.UserId, 10))
	fromData.Add("group_id", strconv.FormatInt(m.GroupId, 10))
	fromData.Add("message", m.Message)
	fromData.Add("auto_escape", fmt.Sprint(m.AutoEscape))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// DeleteMsg 撤回消息
func (h QQHandle) DeleteMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetMsg 获取消息
func (h QQHandle) GetMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetForwardMsg 获取合并转发内容
func (h QQHandle) GetForwardMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_forward_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetImage 获取图片信息
func (h QQHandle) GetImage(file string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("file", file)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_image")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// MarkMsgAsRead 标记消息已读
func (h QQHandle) MarkMsgAsRead(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}

	fromData := make(url.Values)
	fromData.Add("message_id", string(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/mark_msg_as_read")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// SetGroupKick 群组踢人
func (h QQHandle) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) (map[string]interface{}, error) {
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

// SetGroupBan 群组单人禁言
func (h QQHandle) SetGroupBan(groupId int64, userId int64, duration string) (map[string]interface{}, error) {
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

// SetGroupAnonymousBan 群组匿名用户禁言
func (h QQHandle) SetGroupAnonymousBan(groupId int64, flag string, duration string) (map[string]interface{}, error) {
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

// SetGroupWholeBan 群组全员禁言
func (h QQHandle) SetGroupWholeBan(groupId int64, enable bool) (map[string]interface{}, error) {
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

// SetGroupAdmin 群组设置管理员
func (h QQHandle) SetGroupAdmin(groupId int64, userId int64, enable bool) (map[string]interface{}, error) {
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

// SetGroupAnonymous 群组匿名
func (h QQHandle) SetGroupAnonymous(groupId int64, enable bool) (map[string]interface{}, error) {
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

// SetGroupName 修改群聊名称
func (h QQHandle) SetGroupName(groupId int64, groupName string) (map[string]interface{}, error) {
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

// SetGroupLeave 退出群聊
func (h QQHandle) SetGroupLeave(groupId int64, isDismiss bool) (map[string]interface{}, error) {
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

// SetGroupSpecialTitle 设置群组专属头衔
func (h QQHandle) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration string) (map[string]interface{}, error) {
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

// SendGroupSign 群打卡
func (h QQHandle) SendGroupSign(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_group_sign")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// SetFriendAddRequest 处理加好友请求
func (h QQHandle) SetFriendAddRequest(flag string, approve bool, remark string) (map[string]interface{}, error) {
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

// SetGroupAddRequest 处理加群请求／邀请
func (h QQHandle) SetGroupAddRequest(flag string, subType string, approve bool, reason string) (map[string]interface{}, error) {
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

// GetLoginInfo 获取登录号信息
func (h QQHandle) GetLoginInfo() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_login_info"))
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

// QiDianGetAccountInfo 获取企点账号信息
// 该API只有企点协议可用
func (h QQHandle) QiDianGetAccountInfo() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/qidian_get_account_info"))
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

// Profile 账号资料
type Profile struct {
	Nickname     string
	Company      string
	Email        string
	College      string
	PersonalNote string
}

// SetQQProfile 设置登录号资料
func (h QQHandle) SetQQProfile(p Profile) (map[string]interface{}, error) {
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

// GetStrangerInfo 获取陌生人信息
func (h QQHandle) GetStrangerInfo(userId int64, noCache bool) (map[string]interface{}, error) {
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

// GetFriendList 获取好友列表
func (h QQHandle) GetFriendList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_friend_list"))
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

// GetUnidirectionalFriendList 获取单向好友列表
func (h QQHandle) GetUnidirectionalFriendList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_unidirectional_friend_list"))
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

// DeleteFriend 删除好友
func (h QQHandle) DeleteFriend(friendId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("friend_id", strconv.FormatInt(friendId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_friend")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGroupInfo 获取群信息
func (h QQHandle) GetGroupInfo(groupId int64, noCache bool) (map[string]interface{}, error) {
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

// GetGroupList 获取群列表
func (h QQHandle) GetGroupList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_unidirectional_friend_list"))
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

// GetGroupMemberInfo 获取群成员信息
func (h QQHandle) GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (map[string]interface{}, error) {
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

// GetGroupMemberList 获取群成员列表
func (h QQHandle) GetGroupMemberList(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_member_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGroupHonorInfo 获取群荣誉信息
func (h QQHandle) GetGroupHonorInfo(groupId int64, typ string) (map[string]interface{}, error) {
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

// GetCookies 获取 Cookies
// 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
func (h QQHandle) GetCookies(domain string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("domain", domain)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_cookies")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetCsrfToken 获取 CSRF Token
// 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
func (h QQHandle) GetCsrfToken() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_csrf_token"))
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

// GetCredentials 获取 QQ 相关接口凭证
// 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
// 即GetCookies和GetCsrfToken两个接口的合并
func (h QQHandle) GetCredentials(domain string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("domain", domain)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_credentials")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetRecord 获取语音
// 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 提交 pr
// 要使用此接口, 通常需要安装 ffmpeg, 请参考 OneBot 实现的相关说明。
func (h QQHandle) GetRecord(file string, outFormat string) (map[string]interface{}, error) {
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

// CanSendImage 检查是否可以发送图片
func (h QQHandle) CanSendImage() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/can_send_image"))
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

// CanSendRecord 检查是否可以发送语音
func (h QQHandle) CanSendRecord() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/can_send_record"))
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

// GetVersionInfo 获取版本信息
func (h QQHandle) GetVersionInfo() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_version_info"))
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

// SetRestart 重启 go-cqhttp
func (h QQHandle) SetRestart(delay int) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("delay", strconv.Itoa(delay))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_restart")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// CleanCache 清理缓存
func (h QQHandle) CleanCache() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/clean_cache"))
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

// SetGroupPortrait 设置群头像
// 目前这个API在登录一段时间后因cookie失效而失效, 请考虑后使用
func (h QQHandle) SetGroupPortrait(groupId int64, file string, cache int) (map[string]interface{}, error) {
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

// GetWordSlices 获取中文分词 ( 隐藏 API )
// 警告:隐藏 API 是不建议一般用户使用的, 它们只应该在 OneBot 实现内部或由 SDK 和框架使用, 因为不正确的使用可能造成程序运行不正常。
//func (h QQHandle) GetWordSlices(content string) []string {
//	var slices []string
//	fromData := make(url.Values)
//	fromData.Add("content", content)
//	data := strings.NewReader(fromData.Encode())
//	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/.get_word_slices")
//	request, err := http.Post(addr, "application/x-www-form-urlencoded", data)
//	if err != nil {
//		if err!=nil {
//			fmt.Println(err)
//}
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
func (h QQHandle) OcrImage(image string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("image", image)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/ocr_image")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGroupSystemMsg 获取群系统消息
// 如果列表不存在任何消息, 将返回 null
func (h QQHandle) GetGroupSystemMsg() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_system_msg"))
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

// UploadPrivateFile 上传私聊文件
func (h QQHandle) UploadPrivateFile(userId int64, file string, name string) (map[string]interface{}, error) {
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

// UploadGroupFile 上传群文件
func (h QQHandle) UploadGroupFile(groupId int64, file string, name string, folder string) (map[string]interface{}, error) {
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

// GetGroupFileSystemInfo 获取群文件系统信息
func (h QQHandle) GetGroupFileSystemInfo(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_file_system_info")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGroupRootFiles 获取群根目录文件列表
func (h QQHandle) GetGroupRootFiles(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_root_files")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGroupFilesByFolder 获取群子目录文件列表
func (h QQHandle) GetGroupFilesByFolder(groupId int64, folderId string) (map[string]interface{}, error) {
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

// CreateGroupFileFolder 创建群文件文件夹
func (h QQHandle) CreateGroupFileFolder(groupId int64, name string, parentId string) (map[string]interface{}, error) {
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

// DeleteGroupFolder 删除群文件夹
func (h QQHandle) DeleteGroupFolder(groupId int64, folderId string) (map[string]interface{}, error) {
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

// DeleteGroupFile 删除群文件
func (h QQHandle) DeleteGroupFile(groupId int64, fileId string, busId int32) (map[string]interface{}, error) {
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

// GetGroupFileUrl 获取群文件资源链接
func (h QQHandle) GetGroupFileUrl(groupId int64, fileId string, busId int32) (map[string]interface{}, error) {
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

// GetStatus 获取状态
func (h QQHandle) GetStatus() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_status"))
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

// GetGroupAtAllRemain 获取群@全体成员剩余次数
func (h QQHandle) GetGroupAtAllRemain(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_group_at_all_remain")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// HandleQuickOperation 对事件执行快速操作 ( 隐藏 API )
//func (h QQHandle)HandleQuickOperation (){
//
//}

// SendGroupNotice 发送群公告
func (h QQHandle) SendGroupNotice(groupId int64, content string, image string) (map[string]interface{}, error) {
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

// GetGroupNotice 获取群公告
func (h QQHandle) GetGroupNotice(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/_get_group_notice")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// ReloadEventFilter 重载事件过滤器
func (h QQHandle) ReloadEventFilter(file string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("file", file)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/reload_event_filter")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// DownloadFile 下载文件到缓存目录
func (h QQHandle) DownloadFile(fileUrl string, threadCount int32, headers string) (map[string]interface{}, error) {
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

// GetOnlineClients 获取当前账号在线客户端列表
func (h QQHandle) GetOnlineClients(noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("no_cache", fmt.Sprint(noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_online_clients")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGroupMsgHistory 获取群消息历史记录
func (h QQHandle) GetGroupMsgHistory(messageSeq int64, groupId int64) (map[string]interface{}, error) {
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

// SetEssenceMsg 设置精华消息
func (h QQHandle) SetEssenceMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", fmt.Sprint(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_essence_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// DeleteEssenceMsg 移除精华消息
func (h QQHandle) DeleteEssenceMsg(messageId int32) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", fmt.Sprint(messageId))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_essence_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetEssenceMsgList 获取精华消息列表
func (h QQHandle) GetEssenceMsgList(groupId int64) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("group_id", strconv.FormatInt(groupId, 10))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_essence_msg_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// CheckUrlSafely 检查链接安全性
func (h QQHandle) CheckUrlSafely(u string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("url", u)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/check_url_safely")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetModelShow 获取在线机型
func (h QQHandle) GetModelShow(model string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("model", model)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/_get_model_show")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// SetModelShow 设置在线机型
func (h QQHandle) SetModelShow(model string, modelShow string) (map[string]interface{}, error) {
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

//频道

// GetGuildServiceProfile 获取频道系统内BOT的资料
func (h QQHandle) GetGuildServiceProfile() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_service_profile"))
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

// GetGuildList 获取频道列表
func (h QQHandle) GetGuildList() (map[string]interface{}, error) {
	var data map[string]interface{}
	request, err := http.Get(fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_list"))
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

// GetGuildMetaByGuest 通过访客获取频道元数据
func (h QQHandle) GetGuildMetaByGuest(guildID string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_meta_by_guest")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGuildChannelList 获取子频道列表
func (h QQHandle) GetGuildChannelList(guildID string, noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("no_cache", fmt.Sprint(noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_channel_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGuildMemberList 获取频道成员列表
func (h QQHandle) GetGuildMemberList(guildID string, nextToken string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("next_token", fmt.Sprint(nextToken))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_member_list")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGuildMemberProfile 单独获取频道成员信息
func (h QQHandle) GetGuildMemberProfile(guildID string, userID string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("user_id", userID)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_member_profile")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// SendGuildChannelMsg 发送信息到子频道
func (h QQHandle) SendGuildChannelMsg(guildID string, channelID string, message string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("channel_id", channelID)
	fromData.Add("message", message)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/send_guild_channel_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetTopicChannelFeeds 获取话题频道帖子
func (h QQHandle) GetTopicChannelFeeds(guildID string, channelID string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("channel_id", channelID)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_topic_channel_feeds")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// DeleteGuildRole 删除频道角色
func (h QQHandle) DeleteGuildRole(guildID string, roleID string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("role_id", roleID)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/delete_guild_role")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGuildMsg 获取频道消息
func (h QQHandle) GetGuildMsg(messageID string, noCache bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("message_id", messageID)
	fromData.Add("no_cache", fmt.Sprint(noCache))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_msg")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// GetGuildRoles 获取频道角色列表
func (h QQHandle) GetGuildRoles(guildID string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/get_guild_roles")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// SetGuildMemberRole 设置用户在频道中的角色
func (h QQHandle) SetGuildMemberRole(guildID string, set bool, roleID string, users string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("set", fmt.Sprint(set))
	fromData.Add("role_id", roleID)
	fromData.Add("users", users)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/set_guild_member_role")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// UpdateGuildRole 设置用户在频道中的角色
func (h QQHandle) UpdateGuildRole(guildID string, roleID string, name string, color string, indepedent bool) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("role_id", roleID)
	fromData.Add("name", name)
	fromData.Add("color", color)
	fromData.Add("indepedent", fmt.Sprint(indepedent))
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/update_guild_role")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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

// CreateGuildRole 创建频道角色
func (h QQHandle) CreateGuildRole(guildID string, name string, color string, independent bool, initialUsers string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fromData := make(url.Values)
	fromData.Add("guild_id", guildID)
	fromData.Add("color", color)
	fromData.Add("name", name)
	fromData.Add("independent", fmt.Sprint(independent))
	fromData.Add("initial_users", initialUsers)
	readerData := strings.NewReader(fromData.Encode())
	addr := fmt.Sprintf(h.Agreement + "://" + h.Host + "/create_guild_role")
	request, err := http.Post(addr, "application/x-www-form-urlencoded", readerData)
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
