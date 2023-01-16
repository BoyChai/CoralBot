package CoralBot

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

// 是否读取配置文件
var read = false

type ConfigStruct struct {
	Plugin     bool
	PluginInfo bool
}

var Cfg ConfigStruct

// ReadConfig 读取配置文件开关
func ReadConfig() {
	read = true
}

// 检查配置文件是否存在
func checkConfig() error {
	_, err := os.Stat("CoralBot.conf")
	if err == nil {
		return nil
	}
	return err
}

// 创建配置文件
func createConfig() error {
	c, err := os.Create("CoralBot.conf")
	defer func() { _ = c.Close() }()
	if err != nil {
		return err
	}
	return nil
}

// 读取配置文件
func readConfig() error {
	c, err := ini.Load("CoralBot.conf")
	if err != nil {
		return err
	}
	plugin := c.Section("").Key("Plugin").MustBool(true)
	Cfg.Plugin = plugin
	pluginInfo := c.Section("").Key("PluginInfo").MustBool(true)
	Cfg.PluginInfo = pluginInfo
	return nil
}

// 格式化配置文件
func defaultConfig() error {
	c, err := os.OpenFile("CoralBot.conf", os.O_WRONLY, 0644)
	defer func() { _ = c.Close() }()
	if err != nil {
		return err
	}
	defaultConfigData := "# 是否开启插件，默认开启。true or false\nPlugin=true\n# 加载插件时是否输出插件信息，默认开启。true or false\nPluginInfo=true"
	_, err = c.Write([]byte(defaultConfigData))
	if err != nil {
		return err
	}
	return nil
}

// 读取配置文件总流程
func readCoralBotConfig() error {
	err := checkConfig()
	if err != nil {
		fmt.Println("加载主配置文件失败：", err)
		err = nil
		fmt.Println("正在创建并初始化配置文件......")
		err = createConfig()
		if err != nil {
			fmt.Println("创建配置文件失败: ", err)
			return err
		}
		err = defaultConfig()
		if err != nil {
			fmt.Println("初始化配置文件失败: ", err)
			return err
		}
	}
	err = readConfig()
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		return err
	}
	return nil
}
