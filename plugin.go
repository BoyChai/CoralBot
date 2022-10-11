package CoralBot

//PluginInfo 插件信息
type PluginInfo struct {
	// 插件名称
	Name string
	// 插件简介
	Summary string
	// 插件版本
	Version string
	// 插件开发者
	Developer string
	// 插件开发者邮箱
	Email string
	// 当前插件的任务id
	Id int
}

// RequestData 存储Handle返回的值
type RequestData struct {
	Data map[string]interface{}
	Err  error
}

type PluginTool struct {
	H *Handle
	E *Event
}
