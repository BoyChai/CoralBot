package CoralBot

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

// Event 消息全部信息
type Event struct {
	//https://docs.go-cqhttp.org/event/#%E9%80%9A%E7%94%A8%E6%95%B0%E6%8D%AE
	bodyData    string
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	SubType     string `json:"sub_type"`
	RawMessage  string `json:"raw_message"`
	Font        int32  `json:"font"`
	TempSource  int64  `json:"temp_source"`
	Sender      struct {
		Age      int32  `json:"age"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserID   int64  `json:"user_id"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Role     string `json:"role"`
		Title    string `json:"title"`
	} `json:"sender"`
	Anonymous struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	} `json:"anonymous"`
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	MessageSeq int    `json:"message_seq"`
	UserID     int64  `json:"user_id"`
	MessageID  int32  `json:"message_id"`
	NoticeType string `json:"notice_type"`
	File       struct {
		BusID int64  `json:"busid"`
		ID    string `json:"id"`
		Name  string `json:"name"`
		Size  int64  `json:"size"`
		URL   string `json:"url"`
	} `json:"file"`
	OperatorID int64  `json:"operator_id"`
	Duration   int64  `json:"duration"`
	SenderID   int64  `json:"sender_id"`
	TargetID   int64  `json:"target_id"`
	HonorType  string `json:"honor_type"`
	CardNew    string `json:"card_new"`
	CardOld    string `json:"card_old"`
	Comment    string `json:"comment"`
	Flag       string `json:"flag"`
	Client     struct {
		AppId      int64  `json:"app_id"`
		DeviceName string `json:"device_name"`
		DeviceKind string `json:"device_kind"`
	} `json:"client"`
	Online bool `json:"online"`
}

// explain 解析命令函数
func (e *Event) explain() {
	for i := 0; i < cap(Tasks); i++ {
		task := Tasks[i]
		err := json.Unmarshal([]byte(e.bodyData), &e)
		if err != nil {
			fmt.Println("command parsing error,please feedback to the developer.error:", err)
		}
		e.filterStart(task)
	}
}

// 过滤
func (e *Event) filterStart(task Task) {
	for t := 1; t <= cap(task.Condition); t++ {
		//fmt.Println(*task.Condition[t-1].Key.(reflect.TypeOf(task.Condition[t-1].Key))
		conditionKey, _ := e.typeAsserts(task.Condition[t-1].Key)
		if t == cap(task.Condition) {
			if task.Condition[t-1].Regex == true {
				key, _ := regexp.MatchString(task.Condition[t-1].Value, fmt.Sprint(conditionKey))
				if key {
					task.Run()
				}
			}
			if fmt.Sprint(conditionKey) == task.Condition[t-1].Value {
				task.Run()
			}
		}
		if task.Condition[t-1].Regex == true {
			key, _ := regexp.MatchString(task.Condition[t-1].Value, fmt.Sprint(conditionKey))
			if key != true {
				break
			}
		}
		if fmt.Sprint(conditionKey) != task.Condition[t-1].Value {
			break
		}
	}
}

// 类型断言
func (e *Event) typeAsserts(key interface{}) (interface{}, error) {
	switch key.(type) {
	case *int64:
		return *key.(*int64), nil
	case *string:
		return *key.(*string), nil
	default:
		return nil, errors.New("The current type is not supported. Please feedback through issue")
	}
}
