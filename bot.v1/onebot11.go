package botv1

import (
	"time"

	"github.com/tidwall/gjson"
)

type Onebot11Event struct {
	// 消息类型
	MessageType string `json:"message_type"`
	// 消息发起类型
	SubType string `json:"sub_type"`
	// 消息ID
	MessageID int64 `json:"message_id"`
	// 群号
	GroupID int64 `json:"group_id"`
	// 用户QQ号
	UserID int64 `json:"user_id"`
	// 消息原始内容
	Message string `json:"message"`
	// 用户信息
	Sender struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Sex      string `json:"sex"`
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Level    int    `json:"level"`
		Role     string `json:"role"`
		Title    string `json:"title"`
	} `json:"sender"`
	// 时间
	Time time.Time `json:"time"`
	// 机器人自身ID
	SelfID int64 `json:"self_id"`
}

func (e Onebot11Event) GetType() string {
	return "OneBot11"
}

func (e *Onebot11Event) Parse(jsonStr string) {

	// 解析基本字段
	e.MessageType = gjson.Get(jsonStr, "message_type").String()
	e.SubType = gjson.Get(jsonStr, "sub_type").String()
	e.MessageID = gjson.Get(jsonStr, "message_id").Int()
	e.GroupID = gjson.Get(jsonStr, "group_id").Int()
	e.UserID = gjson.Get(jsonStr, "user_id").Int()
	e.Message = gjson.Get(jsonStr, "raw_message").String() // 注意这里使用 raw_message
	e.Time = time.Unix(gjson.Get(jsonStr, "time").Int(), 0)
	e.SelfID = gjson.Get(jsonStr, "self_id").Int()

	// 解析嵌套的 sender 字段
	e.Sender.UserID = gjson.Get(jsonStr, "sender.user_id").Int()
	e.Sender.Nickname = gjson.Get(jsonStr, "sender.nickname").String()
	e.Sender.Card = gjson.Get(jsonStr, "sender.card").String()
	e.Sender.Sex = gjson.Get(jsonStr, "sender.sex").String()
	e.Sender.Age = int(gjson.Get(jsonStr, "sender.age").Int()) // 注意类型转换
	e.Sender.Area = gjson.Get(jsonStr, "sender.area").String()
	e.Sender.Level = int(gjson.Get(jsonStr, "sender.level").Int()) // JSON 中是字符串，需转换
	e.Sender.Role = gjson.Get(jsonStr, "sender.role").String()
	e.Sender.Title = gjson.Get(jsonStr, "sender.title").String()
}

func (e Onebot11Event) Broadcast() error {
	for _, task := range Tasks {
		err := task(&e)
		if err != nil {
			return err
		}
	}
	return nil
}
