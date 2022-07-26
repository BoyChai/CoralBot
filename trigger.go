package CoralBot

type Action struct {
	//触发信息类型
	//https://docs.go-cqhttp.org/reference/data_struct.html#post-type
	//目前支支持message
	Mode    string
	Message string
	Run     func()
}

var AllAction []Action

// NewAction 创建一个动作
func NewAction(mode string, message string, run func()) {
	var action Action
	action.Mode = mode
	action.Message = message
	action.Run = run
	AllAction = append(AllAction, action)
}
