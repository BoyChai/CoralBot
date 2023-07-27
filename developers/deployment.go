package developers

import (
	"errors"
	"fmt"
	"github.com/BoyChai/CoralBot/action"
	"github.com/BoyChai/CoralBot/task"
	"gopkg.in/yaml.v3"
	"os"
)

func init() {
	var err error
	handler, err = readConfig()
	if err != nil {
		fmt.Println("读取插件配置文件出现错误。")
		return
	}
}

func readConfig() (Handler, error) {
	var c pluginConfig
	var all Handler
	var my Handler
	var n Handler
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return n, err
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return Handler{}, err
	}
	for _, value := range c.Handler {
		if value.Name == information.Name {
			my = value
			continue
		}
		if value.Name == "all" {
			all = value
			continue
		}
	}

	if my != n {
		return my, nil
	}
	if all != n {
		return all, nil
	}
	return Handler{}, errors.New("未读取到配置文件")
}

func GetConfigQQHandler() *action.QQHandle {
	return &action.QQHandle{
		Host:      handler.Host,
		Agreement: handler.Agreement,
	}
}

func GetConfigDingDingHandler() *action.DingDingHandle {
	return &action.DingDingHandle{
		AppKey:    handler.AppKey,
		AppSecret: handler.AppSecret,
	}
}

func SetInfo(info task.Plugin) {
	information = info
}

// NewTask 创建插件任务
func NewTask(t task.Task) {
	//task.NewPluginTask(t)
	task.Tasks = append(task.Tasks, t)
}
