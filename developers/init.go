package developers

import (
	"github.com/BoyChai/CoralBot/task"
)

type pluginConfig struct {
	Handler []Handler `yaml:"handler"`
}

type Handler struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Agreement string `yaml:"agreement"`
	AppKey    string `yaml:"appKey"`
	AppSecret string `yaml:"appSecret"`
}

var handler Handler

// 插件的信息
var information task.Plugin
