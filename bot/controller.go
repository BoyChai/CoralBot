package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoyChai/CoralBot/config"
	"github.com/BoyChai/CoralBot/task"
	"github.com/tidwall/gjson"
	"regexp"
	"time"
)

// Explain qq任务解析器
func (e *QQEvent) Explain(bodyData []byte) {
	*e = QQEvent{}
	Tasks := task.Tasks
	err := json.Unmarshal(bodyData, &e)
	if err != nil {
		fmt.Println("command parsing error,please feedback to the developer.error:", err)
	}
	for i := 0; i < len(Tasks); i++ {
		t := Tasks[i]
		if e.MessageType == "guild" {
			e.GuildUserID = gjson.Get(string(bodyData), "user_id").String()
			e.GuildMessageID = gjson.Get(string(bodyData), "message_id").String()
		}
		status := filterStart(t)
		if status == nil {
			// 返回值如果等于空则代此事件已经达成了任务条件并已经执行
			return
		}
	}
}

// Explain DingDing任务解析器
func (e *DingDingEvent) Explain(bodyData []byte) {
	*e = DingDingEvent{}
	err := json.Unmarshal(bodyData, &e)
	if err != nil {
		fmt.Println("command parsing error,please feedback to the developer.error:", err)
	}
	// 获取当前时间戳(毫秒)
	now := time.Now()
	nowTime := now.UnixNano() / 1e6
	// 时间判断是否合法
	if (config.Timestamp-nowTime)/3600000 >= 1 {
		return
	}
	// 判断消息是否合法
	//if config.DingDingSignCheck {
	if config.Cfg.DingDingSignCheck {
		secStr := fmt.Sprintf("%d\n%s", config.Timestamp, config.DingDingAppSecret)
		hmac256 := hmac.New(sha256.New, []byte(config.DingDingAppSecret))
		hmac256.Write([]byte(secStr))
		result := hmac256.Sum(nil)
		sign := base64.StdEncoding.EncodeToString(result)
		if sign != config.Sign {
			fmt.Println("[CoralBot]:机器人识别失败,或签名无效")
			return
		}
	}
	// 任务处理
	Tasks := task.Tasks
	for i := 0; i < len(Tasks); i++ {
		t := Tasks[i]
		status := filterStart(t)
		if status == nil {
			return
		}
	}
}

// 任务执行器
func filterStart(task task.Task) error {
	for t := 1; t <= len(task.Condition); t++ {
		conditionKey, _ := typeAsserts(task.Condition[t-1].Key)
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

// 类型断言
func typeAsserts(key interface{}) (interface{}, error) {
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
