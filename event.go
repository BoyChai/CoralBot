package CoralBot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
)

// Event 消息全部信息
type Event struct {
	//https://docs.go-cqhttp.org/event/
	//bodyData    []byte
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
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
		//频道
		TinyID string `json:"tiny_id"`
	} `json:"sender"`
	Anonymous struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	} `json:"anonymous"`
	GroupID        int64  `json:"group_id"`
	Message        string `json:"message"`
	MessageSeq     int    `json:"message_seq"`
	UserID         int64  `json:"user_id"`
	GuildUserID    string `json:"guild_user_id"`
	MessageID      int32  `json:"message_id"`
	GuildMessageID string `json:"guild_message_id"`
	NoticeType     string `json:"notice_type"`
	MetaEventType  string `json:"meta_event_type"`
	File           struct {
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
		AppID      int64  `json:"app_id"`
		DeviceName string `json:"device_name"`
		DeviceKind string `json:"device_kind"`
	} `json:"client"`
	Online bool `json:"online"`
	//频道
	GuildID          string `json:"guild_id"`
	ChannelID        string `json:"channel_id"`
	CurrentReactions []struct {
		EmojiID    string `json:"emoji_id"`
		EmojiIndex int32  `json:"emoji_index"`
		EmojiType  int32  `json:"emoji_type"`
		EmojiName  string `json:"emoji_name"`
		Count      int32  `json:"count"`
		Clicked    bool   `json:"clicked"`
	} `json:"current_reactions"`
	OldInfo     ChannelInfo `json:"old_info"`
	NewInfo     ChannelInfo `json:"new_info"`
	ChannelInfo ChannelInfo `json:"channel_info"`
	Other       Other
}

// ChannelInfo 频道信息
type ChannelInfo struct {
	OwnerGuildID    string `json:"owner_guild_id"`
	ChannelID       string `json:"channel_id"`
	ChannelType     int32  `json:"channel_type"`
	ChannelName     string `json:"channel_name"`
	CreateTime      int64  `json:"create_time"`
	CreatorTinyID   string `json:"creator_tiny_id"`
	TalkPermission  int32  `json:"talk_permission"`
	VisibleType     int32  `json:"visible_type"`
	CurrentSlowMode int32  `json:"current_slow_mode"`
	SlowModes       []struct {
		SlowModeKey    int32  `json:"slow_mode_key"`
		SlowModeText   string `json:"slow_mode_text"`
		SpeakFrequency int32  `json:"speak_frequency"`
		SlowModeCircle int32  `json:"slow_mode_circle"`
	} `json:"slow_modes"`
}

// Other 其他
type Other struct {
	// 运行插件
	RunName string
}

// explain 解析命令函数
// 将任务以此拿出来进行判断
func (e *Event) explain(bodyData []byte) {
	for i := 0; i < len(Tasks); i++ {
		task := Tasks[i]
		err := json.Unmarshal(bodyData, &e)
		if e.MessageType == "guild" {
			e.GuildUserID = gjson.Get(string(bodyData), "user_id").String()
			e.GuildMessageID = gjson.Get(string(bodyData), "message_id").String()
		}
		if err != nil {
			fmt.Println("command parsing error,please feedback to the developer.error:", err)
		}
		// 判断是否为插件
		if task.Plugin {
			status := e.pluginFilterStart(task)
			// 返回值如果等于空则代此事件已经达成了任务条件并已经执行
			if status == nil {
				return
			}
		} else {
			status := e.filterStart(task)
			if status == nil {
				// 返回值如果等于空则代此事件已经达成了任务条件并已经执行
				return
			}
		}

	}
}

// 过滤
func (e *Event) filterStart(task Task) error {
	for t := 1; t <= len(task.Condition); t++ {
		conditionKey, _ := e.typeAsserts(task.Condition[t-1].Key)
		// 如果这是此任务的最后一个判断
		if t == len(task.Condition) {
			if task.Condition[t-1].Regex == true {
				// 正则判断
				key, _ := regexp.MatchString(task.Condition[t-1].Value, fmt.Sprint(conditionKey))
				if key {
					task.Run()
					return nil
				}
			}
			if fmt.Sprint(conditionKey) == task.Condition[t-1].Value {
				task.Run()
				return nil
			}
		}
		if task.Condition[t-1].Regex == true {
			key, _ := regexp.MatchString(task.Condition[t-1].Value, fmt.Sprint(conditionKey))
			if key != true {
				return errors.New("1")
			}
		}
		if fmt.Sprint(conditionKey) != task.Condition[t-1].Value {
			return errors.New("1")
		}
	}
	return errors.New("1")
}

func (e *Event) pluginFilterStart(task Task) error {
	var pts []Task
	//fmt.Println("错误信息为", task.pingoServer)
	err := task.pingoServer.Call("Plugin.GetPlugin", &e, &pts)
	if err != nil {
		fmt.Println("插件任务获取失败:插件名称", task.Info.Name, "\n", "错误信息为:", err)
	} else {
		// 判断插件条件
		// 依次判断任务
		for ti := 0; ti < len(pts); ti++ {
			// 依次判断任务条件
			for ci := 0; ci < len(pts[ti].Condition); ci++ {
				if ci+1 == len(pts[ti].Condition) {
					// 正则判断
					if pts[ti].Condition[ci].Regex == true {
						key, _ := regexp.MatchString(pts[ti].Condition[ci].Value, fmt.Sprint(pts[ti].Condition[ci].Key))
						if key {
							e.Other.RunName = pts[ti].RunName
							err := task.pingoServer.Call("Plugin.Handles", &e, nil)
							e.Other.RunName = ""
							if err != nil {
								fmt.Println("插件任务调用错误:", err)
							}
							return nil
						}
					}
					if fmt.Sprint(pts[ti].Condition[ci].Key) == pts[ti].Condition[ci].Value {
						e.Other.RunName = pts[ti].RunName
						err := task.pingoServer.Call("Plugin.Handles", &e, nil)
						e.Other.RunName = ""
						if err != nil {
							fmt.Println("插件任务调用错误:", err)
						}
						return nil
					}
				}
				if pts[ti].Condition[ci].Regex == true {
					key, _ := regexp.MatchString(pts[ti].Condition[ci].Value, fmt.Sprint(pts[ti].Condition[ci].Key))
					if key != true {
						return errors.New("1")
					}
				}
				if fmt.Sprint(pts[ti].Condition[ci].Key) != pts[ti].Condition[ci].Value {
					return errors.New("1")
				}
			}
		}
		return errors.New("1")
	}
	return errors.New("1")
}

// 类型断言
func (e *Event) typeAsserts(key interface{}) (interface{}, error) {
	switch key.(type) {
	case *int64:
		return *key.(*int64), nil
	case *string:
		return *key.(*string), nil
	case *int32:
		return *key.(*int32), nil
	default:
		return nil, errors.New("the current type is not supported. please feedback through issue")
	}
}
