package task

import "net"

// Plugin 插件信息
type Plugin struct {
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
	// 插件的连接
	accept net.Conn
}

// SetAccept 设置插件连接
func (p *Plugin) SetAccept(conn net.Conn) {
	p.accept = conn
}

// GetAccept 获取插件连接
func (p *Plugin) GetAccept() net.Conn {
	return p.accept
}

type PluginType string

const QQ PluginType = "QQ"
const DingDing PluginType = "DingDing"
