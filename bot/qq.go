package bot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type QQEvent struct {
	//https://docs.go-cqhttp.org/event/
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

func (e QQEvent) GetType() string {
	return "QQ"
}

func (e QQEvent) GetDocs() string {
	return "https://docs.go-cqhttp.org/"
}
func (e *QQEvent) GetRunName() string {
	return e.Other.RunName
}
func (e *QQEvent) SetRunName(runName string) {
	e.Other.RunName = runName
}

func (e *QQEvent) GetLogOut(params gin.LogFormatterParams) string {
	logout := "[CoralBot] QQBot:"
	logout += fmt.Sprintf(" 时间: %s 事件内容: %+v\n",
		params.TimeStamp.Format(time.RFC3339),
		e)
	return logout
}
