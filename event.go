package CoralBot

import (
	"fmt"
	"regexp"

	"github.com/tidwall/gjson"
)

type Event struct {
	//https://docs.go-cqhttp.org/event/#%E9%80%9A%E7%94%A8%E6%95%B0%E6%8D%AE
	bodyData    string
	Time        string
	SelfID      string
	PostType    string
	MessageType string
	SubType     string
	MessageId   string
	UserID      string
	Message     string
	RawMessage  string
	GroupID     string
	Font        string
	Sender      Sender
	TempSource  string
	Anonymous   Anonymous
	NoticeType  string
	File        File
	OperatorId  string
	Duration    string
	TargetId    string
	Comment     string
	Flag        string
}
type Sender struct {
	UserID   string
	Nickname string
	Sex      string
	Age      string
	Card     string
	Area     string
	Level    string
	Role     string
	Title    string
}
type Anonymous struct {
	ID   string
	Name string
	Flag string
}
type File struct {
	ID    string
	Name  string
	Size  string
	BusId string
}

//func (e *Event) Parse(boyData []byte) {
//e.Time = gjson.Get(string(boyData), "time").String()
//e.SelfID = gjson.Get(string(boyData), "self_id").String()
//e.PostType = gjson.Get(string(boyData), "post_type").String()
//e.SubType = gjson.Get(string(boyData), "sub_type").String()
//e.MessageId = gjson.Get(string(boyData), "message_id").String()
//e.UserID = gjson.Get(string(boyData), "user_id").String()
//e.GroupID = gjson.Get(string(boyData), "group_id").String()
//e.Message = gjson.Get(string(boyData), "message").String()
//e.RawMessage = gjson.Get(string(boyData), "raw_message").String()
//e.Font = gjson.Get(string(boyData), "font").String()
//e.Sender.UserID = gjson.Get(string(boyData), "sender.user_id").String()
//e.Sender.Nickname = gjson.Get(string(boyData), "sender.nickname").String()
//e.Sender.Sex = gjson.Get(string(boyData), "sender.sex").String()
//e.Sender.Age = gjson.Get(string(boyData), "sender.age").String()
//e.Sender.Card = gjson.Get(string(boyData), "sender.card").String()
//e.Sender.Area = gjson.Get(string(boyData), "sender.area").String()
//e.Sender.Level = gjson.Get(string(boyData), "sender.level").String()
//e.Sender.Role = gjson.Get(string(boyData), "sender.role").String()
//e.Sender.Title = gjson.Get(string(boyData), "sender.title").String()
//e.TempSource = gjson.Get(string(boyData), "temp_source").String()
//e.Anonymous.ID = gjson.Get(string(boyData), "anonymous.id").String()
//e.Anonymous.Name = gjson.Get(string(boyData), "anonymous.name").String()
//e.Anonymous.Flag = gjson.Get(string(boyData), "anonymous.flag").String()
//e.NoticeType = gjson.Get(string(boyData), "notice_type").String()
//e.File.ID = gjson.Get(string(boyData), "file.id").String()
//e.File.Name = gjson.Get(string(boyData), "file.name").String()
//e.File.Size = gjson.Get(string(boyData), "file.size").String()
//e.File.BusId = gjson.Get(string(boyData), "file.busid").String()
//e.OperatorId = gjson.Get(string(boyData), "operator_id").String()
//e.Duration = gjson.Get(string(boyData), "duration").String()
//e.TargetId = gjson.Get(string(boyData), "target_id").String()
//e.Comment = gjson.Get(string(boyData), "comment").String()
//e.Flag = gjson.Get(string(boyData), "flag").String()
//}

func (e *Event) all_message() {
	e.Time = gjson.Get(e.bodyData, "time").String()
	e.SelfID = gjson.Get(e.bodyData, "self_id").String()
	e.PostType = gjson.Get(e.bodyData, "post_type").String()
	e.MessageType = gjson.Get(e.bodyData, "message_type").String()
	e.SubType = gjson.Get(e.bodyData, "sub_type").String()
	e.MessageId = gjson.Get(e.bodyData, "message_id").String()
	e.UserID = gjson.Get(e.bodyData, "user_id").String()
	e.GroupID = gjson.Get(e.bodyData, "group_id").String()
	e.Message = gjson.Get(e.bodyData, "message").String()
	e.RawMessage = gjson.Get(e.bodyData, "raw_message").String()
	e.Font = gjson.Get(e.bodyData, "font").String()
	e.Sender.UserID = gjson.Get(e.bodyData, "sender.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "sender.nickname").String()
	e.Sender.Sex = gjson.Get(e.bodyData, "sender.sex").String()
	e.Sender.Age = gjson.Get(e.bodyData, "sender.age").String()
	e.Sender.Card = gjson.Get(e.bodyData, "sender.card").String()
	e.Sender.Area = gjson.Get(e.bodyData, "sender.area").String()
	e.Sender.Level = gjson.Get(e.bodyData, "sender.level").String()
	e.Sender.Role = gjson.Get(e.bodyData, "sender.role").String()
	e.Sender.Title = gjson.Get(e.bodyData, "sender.title").String()
	e.TempSource = gjson.Get(e.bodyData, "temp_source").String()
	e.Anonymous.ID = gjson.Get(e.bodyData, "anonymous.id").String()
	e.Anonymous.Name = gjson.Get(e.bodyData, "anonymous.name").String()
	e.Anonymous.Flag = gjson.Get(e.bodyData, "anonymous.flag").String()
}

func (e *Event) private_message() {
	e.Time = gjson.Get(e.bodyData, "time").String()
	e.SelfID = gjson.Get(e.bodyData, "self_id").String()
	e.PostType = gjson.Get(e.bodyData, "post_type").String()
	e.MessageType = gjson.Get(e.bodyData, "message_type").String()
	e.SubType = gjson.Get(e.bodyData, "sub_type").String()
	e.MessageId = gjson.Get(e.bodyData, "message_id").String()
	e.UserID = gjson.Get(e.bodyData, "user_id").String()
	e.Message = gjson.Get(e.bodyData, "message").String()
	e.RawMessage = gjson.Get(e.bodyData, "raw_message").String()
	e.Font = gjson.Get(e.bodyData, "font").String()
	e.Sender.UserID = gjson.Get(e.bodyData, "sender.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "sender.nickname").String()
	e.Sender.Sex = gjson.Get(e.bodyData, "sender.sex").String()
	e.Sender.Age = gjson.Get(e.bodyData, "sender.age").String()
	e.Sender.Card = gjson.Get(e.bodyData, "sender.card").String()
	e.Sender.Area = gjson.Get(e.bodyData, "sender.area").String()
	e.Sender.Level = gjson.Get(e.bodyData, "sender.level").String()
	e.Sender.Role = gjson.Get(e.bodyData, "sender.role").String()
	e.Sender.Title = gjson.Get(e.bodyData, "sender.title").String()
	e.TempSource = gjson.Get(e.bodyData, "temp_source").String()
}

func (e *Event) group_message() {
	e.Time = gjson.Get(e.bodyData, "time").String()
	e.SelfID = gjson.Get(e.bodyData, "self_id").String()
	e.PostType = gjson.Get(e.bodyData, "post_type").String()
	e.MessageType = gjson.Get(e.bodyData, "message_type").String()
	e.SubType = gjson.Get(e.bodyData, "sub_type").String()
	e.MessageId = gjson.Get(e.bodyData, "message_id").String()
	e.UserID = gjson.Get(e.bodyData, "user_id").String()
	e.GroupID = gjson.Get(e.bodyData, "group_id").String()
	e.Message = gjson.Get(e.bodyData, "message").String()
	e.RawMessage = gjson.Get(e.bodyData, "raw_message").String()
	e.Font = gjson.Get(e.bodyData, "font").String()
	e.Sender.UserID = gjson.Get(e.bodyData, "sender.user_id").String()
	e.Sender.Nickname = gjson.Get(e.bodyData, "sender.nickname").String()
	e.Sender.Sex = gjson.Get(e.bodyData, "sender.sex").String()
	e.Sender.Age = gjson.Get(e.bodyData, "sender.age").String()
	e.Sender.Card = gjson.Get(e.bodyData, "sender.card").String()
	e.Sender.Area = gjson.Get(e.bodyData, "sender.area").String()
	e.Sender.Level = gjson.Get(e.bodyData, "sender.level").String()
	e.Sender.Role = gjson.Get(e.bodyData, "sender.role").String()
	e.Sender.Title = gjson.Get(e.bodyData, "sender.title").String()
	e.Anonymous.ID = gjson.Get(e.bodyData, "anonymous.id").String()
	e.Anonymous.Name = gjson.Get(e.bodyData, "anonymous.name").String()
	e.Anonymous.Flag = gjson.Get(e.bodyData, "anonymous.flag").String()
}

// explain 解析命令函数
func (e *Event) explain() {
	for i := 0; i < cap(Tasks); i++ {
		task := Tasks[i]
		switch task.Mode {
		case "all_message":
			e.all_message()
			for t := 0; t < cap(task.Condition); t++ {
				if task.Condition[t].Regex == true {
					key, _ := regexp.MatchString(*task.Condition[t].Key, task.Condition[t].Value)
					if key {
						task.Run()
					}
					return
				}
				if *task.Condition[t].Key == task.Condition[t].Value {
					task.Run()
				}
			}
		case "private_message":
			e.private_message()
			for t := 0; t < cap(task.Condition); t++ {
				if task.Condition[t].Regex == true {
					key, _ := regexp.MatchString(*task.Condition[t].Key, task.Condition[t].Value)
					if key {
						task.Run()
					}
					return
				}
				if *task.Condition[t].Key == task.Condition[t].Value {
					task.Run()
				}
			}
		case "group_message":
			e.group_message()
			for t := 0; t < cap(task.Condition); t++ {
				if task.Condition[t].Regex == true {
					key, _ := regexp.MatchString(*task.Condition[t].Key, task.Condition[t].Value)
					if key {
						task.Run()
					}
					return
				}
				if *task.Condition[t].Key == task.Condition[t].Value {
					task.Run()
				}
			}
		default:
			fmt.Printf("%v事件解析失败", task.Mode)
		}
	}
}
