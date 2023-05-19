package developers

import (
	"github.com/BoyChai/CoralBot/task"
)

//type Handle struct {
//	name string
//	//run  func(event *bot.Event) error
//	//run func(event *bot.QQEvent) error
//	run func(event interface{}) error
//}

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

// var event bot.Event
// var event bot.QQEvent
//var event interface{}

//var qqEvent bot.QQEvent

//var dingDingEvent bot.DingDingEvent

//var qqHan action.QQHandle

//var dingDingHan action.DingDingHandle

var information task.Info

//type Plugin struct{}

//var handles []Handle
