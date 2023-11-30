package structure

// QQMsg QQ消息
type QQMsg struct {
	UserId     int64
	GroupId    int64
	Message    string
	AutoEscape bool
}

// Profile 账号资料
type Profile struct {
	Nickname     string
	Company      string
	Email        string
	College      string
	PersonalNote string
}
