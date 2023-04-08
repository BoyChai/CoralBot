package bot

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
}

func (e DingDingEvent) GetType() string {
	return "DingDing"
}

func (e DingDingEvent) GetDocs() string {
	return "https://open.dingtalk.com/document/orgapp/robot-overview"
}
