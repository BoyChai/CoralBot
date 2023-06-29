package config

import (
	"fmt"
	"os"
)
import "github.com/go-ini/ini"

type ConfigStruct struct {
	Listen            int
	Plugin            bool
	PluginInfo        bool
	DingDingSignCheck bool
}

var Cfg ConfigStruct

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
	Cfg.Plugin = c.Section("").Key("plugin").MustBool(true)
	Cfg.PluginInfo = c.Section("").Key("plugin_info").MustBool(true)
	Cfg.Listen = c.Section("").Key("listen").MustInt(8080)
	Cfg.DingDingSignCheck = c.Section("").Key("dingding_sign_check").MustBool(true)
	return nil
}

// 格式化配置文件
func defaultConfig() error {
	c, err := os.OpenFile("CoralBot.conf", os.O_WRONLY, 0644)
	defer func() { _ = c.Close() }()
	if err != nil {
		return err
	}
	defaultConfigData := "# CoralBot监听端口，默认为8080。\nlisten=8080\n# 是否开启插件，默认开启。true or false\nplugin=true\n# 加载插件时是否输出插件信息，默认开启。true or false\nplugin_info=true\n# 钉钉是否检查上报信息的合法性，默认为true\ndingding_sign_check=true"
	_, err = c.Write([]byte(defaultConfigData))
	if err != nil {
		return err
	}
	return nil
}

// ReadCoralBotConfig 读取配置文件总流程
func ReadCoralBotConfig() error {
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
