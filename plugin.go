package CoralBot

import (
	"github.com/dullgiulio/pingo"
	"io/ioutil"
	"strings"
)

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

// 读取插件
func loadPlugin() error {
	// 读插件
	files, err := ioutil.ReadDir("./plugin")
	if err != nil {
		return err
	}
	for _, file := range files {
		// 识别插件
		if strings.HasSuffix(file.Name(), ".coral") {
			// 创建pingoServer
			strings.Split(file.Name(), ".")
			pingoServer := pingo.NewPlugin("tcp", "./plugin/"+file.Name())
			pingoServer.Start()
			defer pingoServer.Stop()
			// 读取插件信息
			var info PluginInfo
			err := pingoServer.Call("MyPlugin.PluginInfo", "", &info)
			if err != nil {
				return err
			}
			// 读取插件任务
			var t []Task
			err = pingoServer.Call("MyPlugin.GetPlugin", "", &t)
			if err != nil {
				return err
			}
			for i := 0; i < len(t); i++ {
				// 将信息加载到创建任务里面
				t[i].info.Name = info.Name
				t[i].info.Developer = info.Developer
				t[i].info.Email = info.Email
				t[i].info.Version = info.Version
				t[i].info.Summary = info.Summary
				t[i].info.Id = i
				// 设置pingoServer
				t[i].pingoServer = pingoServer
				// 将任务加载到本地
				Tasks = append(Tasks, t[i])
			}
		}
		return nil
	}
	return nil
}
