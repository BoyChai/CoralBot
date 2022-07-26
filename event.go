package CoralBot

import (
	"github.com/tidwall/gjson"
)

type Event struct {
	//https://docs.go-cqhttp.org/event/#%E9%80%9A%E7%94%A8%E6%95%B0%E6%8D%AE
	Time       string
	SelfID     string
	PostType   string
	SubType    string
	MessageId  string
	UserID     string
	Message    string
	RawMessage string
	GroupID    string
	Font       string
	//sender    string
}

func (e *Event) Parse(boyData []byte) {
	e.Time = gjson.Get(string(boyData), "time").String()
	e.SelfID = gjson.Get(string(boyData), "self_id").String()
	e.PostType = gjson.Get(string(boyData), "post_type").String()
	e.SubType = gjson.Get(string(boyData), "sub_type").String()
	e.MessageId = gjson.Get(string(boyData), "message_id").String()
	e.UserID = gjson.Get(string(boyData), "user_id").String()
	e.GroupID = gjson.Get(string(boyData), "group_id").String()
	e.Message = gjson.Get(string(boyData), "message").String()
	e.RawMessage = gjson.Get(string(boyData), "raw_message").String()
	e.Font = gjson.Get(string(boyData), "font").String()
	e.explain()
}

// explain 解析命令函数
func (e *Event) explain() {
	for i := 0; i < cap(AllAction); i++ {
		action := AllAction[i]
		if e.PostType == action.Mode {
			if e.Message == action.Message {
				action.Run()
			}
		}
	}
}
