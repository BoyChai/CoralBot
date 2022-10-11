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
	// 执行调用函数
	RunName string
}

// RequestData 存储Handle返回的值
type RequestData struct {
	data map[string]interface{}
	err  error
}
