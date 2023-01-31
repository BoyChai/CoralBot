package CoralBot

import (
	"fmt"
	"github.com/dullgiulio/pingo"
	"io/ioutil"
	"strings"
)

// PluginInfo 插件信息
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
}

type PluginTool struct {
	H *Handle
	E *Event
}

// 读取插件
func (e *Event) loadPlugin() error {
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
			// 读取插件信息
			var info PluginInfo
			err := pingoServer.Call("Plugin.PluginInfo", "", &info)
			if err != nil {
				return err
			}
			// 加载插件到本地
			var t Task
			t.Info = info
			t.pingoServer = pingoServer
			t.Plugin = true
			Tasks = append(Tasks, t)
		}
	}
	if Cfg.PluginInfo {
		sum := 0
		for i := 0; i < len(Tasks); i++ {
			if Tasks[i].Plugin {
				fmt.Println("Loading succeeded：", Tasks[i].Info.Name)
				fmt.Println("===============Plugin-Info===============")
				fmt.Println("插件名称：", Tasks[i].Info.Name)
				fmt.Println("插件版本：", Tasks[i].Info.Version)
				fmt.Println("插件概述：", Tasks[i].Info.Summary)
				fmt.Println("插件作者：", Tasks[i].Info.Developer)
				fmt.Println("作者邮箱：", Tasks[i].Info.Email)
				sum++
			}
			fmt.Println("=========================================")
			fmt.Println("CoralBot加载插件数量为：", sum)
		}
	} else {
		sum := 0
		for i := 0; i < len(Tasks); i++ {
			if Tasks[i].Plugin {
				fmt.Println("Loading succeeded:", Tasks[i].Info.Name)
				sum++
			}
		}
		fmt.Println("CoralBot加载插件数量为：", sum)
	}
	return nil
}
