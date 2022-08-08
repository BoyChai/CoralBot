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

// GroupInfo 群聊详细信息
// 此结构体主要用来返回值
type GroupInfo struct {
	GroupId         string
	GroupName       string
	GroupMemo       string
	GroupCreateTime string
	GroupLevel      string
	MemberCount     string
	MaxMemberCount  string
}

// GroupHonor 群荣誉信息
// 此结构体主要用来返回值
type GroupHonor struct {
	GroupId          string
	CurrentTalkative Honor
	TalkativeList    []Honor
	PerformerList    []Honor
	LegendList       []Honor
	StrongNewbieList []Honor
	EmotionList      []Honor
}

// Honor 荣誉的具体信息
// 此结构体主要用来返回值
type Honor struct {
	UserId      string
	Nickname    string
	Avatar      string
	DayCount    string
	Description string
}

// AppInfo go-cqhttp版本信息
// 此结构体主要用来返回值
type AppInfo struct {
	AppName                  string
	AppVersion               string
	AppFullName              string
	ProtocolVersion          string
	CoolQEdition             string
	CoolQDirectory           string
	GoCqHttp                 string
	PluginVersion            string
	PluginBuildNumber        string
	PluginBuildConfiguration string
	RuntimeVersion           string
	RuntimeOs                string
	Version                  string
	Protocol                 string
}

// ImageOCR 图片 OCR返回数据
type ImageOCR struct {
	Texts    []TextDetection
	Language string
}

// TextDetection 用于图片OCR结构体
type TextDetection struct {
	Test        string
	Confidence  string
	Coordinates string
}

// InvitedRequest 用于群系统邀请消息列表
type InvitedRequest struct {
	RequestId   string
	InvitorUin  string
	InvitorNick string
	GroupId     string
	GroupName   string
	Checked     string
	Actor       string
}

// JoinRequest 用于群系统邀请消息列表
type JoinRequest struct {
	RequestId     string
	RequesterUin  string
	RequesterNick string
	Message       string
	GroupId       string
	GroupName     string
	Checked       string
	Actor         string
}

// GroupFileSystemInfo 群文件信息用于群文件相关函数Handle返回
type GroupFileSystemInfo struct {
	FileCount  string
	LimitCount string
	UsedSpace  string
	TotalSpace string
}

// GroupFile 群文件信息用于群文件相关函数Handle返回
type GroupFile struct {
	GroupId       string
	FileId        string
	FileName      string
	BusId         string
	FileSize      string
	UploadTime    string
	DeadTime      string
	ModifyTime    string
	DownloadTimes string
	Uploader      string
	UploaderName  string
}

// Folder 群文件夹信息用于群文件相关函数Handle返回
type Folder struct {
	GroupId        string
	FolderId       string
	FolderName     string
	CreateTime     string
	Creator        string
	CreatorName    string
	TotalFileCount string
}

// Status 当前状态结构体 用于api返回
type Status struct {
	AppInitialized string
	AppEnabled     string
	PluginsGood    string
	AppGood        string
	Online         string
	Good           string
	Stat           Statistics
}

// Statistics 当前状态结构体 用于api返回
type Statistics struct {
	PacketReceived  string
	PacketSent      string
	PacketLost      string
	MessageReceived string
	MessageSent     string
	DisconnectTimes string
	LostTimes       string
	LastMessageTime string
}

// GroupAtAllRemain @全体成员剩余次数 用于api返回
type GroupAtAllRemain struct {
	CanAtAll                 string
	RemainAtAllCountForGroup string
	RemainAtAllCountForUin   string
}

// GroupNotice 群公告信息 用于api返回
type GroupNotice struct {
	SenderId    string
	PublishTime string
	Message     GroupNoticeMessage
}
type GroupNoticeMessage struct {
	Text   string
	Images GroupNoticeImages
}
type GroupNoticeImages struct {
	Height string
	Width  string
	Id     string
}

// Clients 账号在线客户端数据 用于api返回
type Clients struct {
	AppId      string
	DeviceName string
	DeviceKind string
}

// EssenceMsg 精华消息信息 用于api返回
type EssenceMsg struct {
	SenderId     string
	SenderNick   string
	SenderTime   string
	OperatorId   string
	OperatorNick string
	OperatorTime string
	MessageId    string
}

// Variant 获取在线机型 用于api返回
type Variant struct {
	ModelShow string
	NeedPay   string
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
