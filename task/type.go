package task

// Info 插件信息
type Info struct {
	// 插件类型
	Type PluginType
	//Type string
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
}

type PluginType string

const QQ PluginType = "QQ"
const DingDing PluginType = "DingDing"
