package CoralBot

import (
	"fmt"
	"regexp"

	"github.com/tidwall/gjson"
)

// Event 消息全部信息
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
	RealId      string
	Messages    []ForwardMessage
	Image       Image
	QiDian      QiDian
}

// Sender 账号信息
type Sender struct {
	GroupID         string
	UserID          string
	Nickname        string
	Sex             string
	Age             string
	QID             string
	Card            string
	Area            string
	Level           string
	Role            string
	Title           string
	LoginDays       string
	Remark          string
	Source          string
	JoinTime        string
	LastSentTime    string
	Unfriendly      string
	TitleExpireTime string
	CardChangeable  string
	ShutUpTimestamp string
}

// Anonymous 匿名信息
type Anonymous struct {
	ID   string
	Name string
	Flag string
}

// File 文件信息
type File struct {
	ID    string
	Name  string
	Size  string
	BusId string
}

// ForwardMessage 合并转发内容
type ForwardMessage struct {
	Content string
	Sender  Sender
	Time    string
}

// Image 图片信息
type Image struct {
	Size     string
	Filename string
	Url      string
}

// QiDian 企点资料
type QiDian struct {
	MasterId   string
	ExtName    string
	CreateTime string
}

// Profile 账号资料
// 此结构体主要用来api传参
type Profile struct {
	Nickname     string
	Company      string
	Email        string
	College      string
	PersonalNote string
}

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
		// switch判断触发器类型
		switch task.Mode {
		case "all_message":
			e.all_message()
			e.filterStart(task)
		case "private_message":
			e.private_message()
			e.filterStart(task)
		case "group_message":
			e.group_message()
			e.filterStart(task)
		default:
			fmt.Printf("%v事件解析失败", task.Mode)
		}
	}
}
func (e *Event) filterStart(task Task) {
	for t := 1; t <= cap(task.Condition); t++ {
		if t == cap(task.Condition) {
			if task.Condition[t-1].Regex == true {
				key, _ := regexp.MatchString(task.Condition[t-1].Value, *task.Condition[t-1].Key)
				if key {
					task.Run()
				}
			}
			if *task.Condition[t-1].Key == task.Condition[t-1].Value {
				task.Run()
			}
		}
		if task.Condition[t-1].Regex == true {
			key, _ := regexp.MatchString(task.Condition[t-1].Value, *task.Condition[t-1].Key)
			if key != true {
				break
			}
		}
		if *task.Condition[t-1].Key != task.Condition[t-1].Value {
			break
		}
	}
}
