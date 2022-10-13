package main

import (
	"github.com/BoyChai/CoralBot"
	"github.com/dullgiulio/pingo"
)

var H CoralBot.Handle

var Info CoralBot.PluginInfo

type MyPlugin struct{}

func (p *MyPlugin) PluginInfo(n string, info *CoralBot.PluginInfo) error {
	*info = Info
	return nil
}

func (p *MyPlugin) GetPlugin(e *CoralBot.Event, Task *[]CoralBot.Task) error {
	CoralBot.NewPluginTask(CoralBot.Task{
		Condition: []CoralBot.Condition{
			{
				Key:   &e.Message,
				Value: "插件测试",
			},
		},
		RunName: "MyPlugin.Run1",
	})
	*Task = CoralBot.Tasks
	return nil
}

func (p *MyPlugin) Run1(e CoralBot.Event, n *string) error {
	h := CoralBot.Handle{
		Host:      "127.0.0.1:5700",
		Agreement: "http",
	}
	h.Reply(e, CoralBot.Msg{
		Message: "插件激活成功",
	})
	return nil
}
func main() {
	H.Host = "127.0.0.1:5700"
	H.Agreement = "http"
	Info = CoralBot.PluginInfo{
		Name:      "test_demo",
		Summary:   "测试插件demo",
		Version:   "v0.0.1",
		Developer: "BoyChai",
		Email:     "1972567225@qq.com",
	}

	plugin := &MyPlugin{}

	pingo.Register(plugin)
	pingo.Run()
}
