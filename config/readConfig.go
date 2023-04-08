package config

import (
	"fmt"
	"os"
)
import "github.com/go-ini/ini"

type ConfigStruct struct {
	Listen int
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
	listen := c.Section("").Key("listen").MustInt(8080)
	Cfg.Listen = listen
	return nil
}

// 格式化配置文件
func defaultConfig() error {
	c, err := os.OpenFile("CoralBot.conf", os.O_WRONLY, 0644)
	defer func() { _ = c.Close() }()
	if err != nil {
		return err
	}
	defaultConfigData := "# CoralBot监听端口，默认为8080。\nlisten=8080\n"
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
