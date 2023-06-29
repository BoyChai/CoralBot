package bot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// DingDingEvent 整个事件信息
type DingDingEvent struct {
	HandleID       int
	ConversationId string `json:"conversationId"`
	AtUsers        []struct {
		DingtalkId string `json:"dingtalkId"`
		StaffId    string `json:"staffId"`
	} `json:"atUsers"`
	ChatbotCorpId             string `json:"chatbotCorpId"`
	ChatbotUserId             string `json:"chatbotUserId"`
	MsgId                     string `json:"msgId"`
	SenderNick                string `json:"senderNick"`
	IsAdmin                   bool   `json:"isAdmin"`
	SenderStaffId             string `json:"senderStaffId"`
	SessionWebhookExpiredTime int64  `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64  `json:"createAt"`
	SenderCorpId              string `json:"senderCorpId"`
	ConversationType          string `json:"conversationType"`
	SenderId                  string `json:"senderId"`
	ConversationTitle         string `json:"conversationTitle"`
	IsInAtList                bool   `json:"isInAtList"`
	SessionWebhook            string `json:"sessionWebhook"`
	Text                      struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
	Other   Other
}

func (e DingDingEvent) GetType() string {
	return "DingDing"
}

func (e DingDingEvent) GetDocs() string {
	return "https://open.dingtalk.com/document/orgapp/robot-overview"
}

func (e *DingDingEvent) GetRunName() string {
	return e.Other.RunName
}

func (e *DingDingEvent) SetRunName(runName string) {
	e.Other.RunName = runName
}

func (e *DingDingEvent) GetLogOut(params gin.LogFormatterParams) string {
	logout := "[CoralBot] DingDingBot:"
	logout += fmt.Sprintf(" 时间: %s 事件内容: %+v\n",
		params.TimeStamp.Format(time.RFC3339),
		e)
	return logout
}
