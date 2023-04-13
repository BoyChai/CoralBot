package developers

import (
	"encoding/gob"
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/task"
	"net"
	"os"
)

func RunPlugin() {
	//// 删除已经存在的套接字文件
	if err := os.RemoveAll(getSocketName(true)); err != nil {
		fmt.Println("flow: 1;启动插件出现错误:", err)
		return
	}
	//// 监听 Unix 域套接字
	listener, err := net.Listen("unix", getSocketName(true))
	if err != nil {
		fmt.Println("flow: 2;启动插件出现错误:", err)
		return
	}

	// 第一次发包 声明我是谁
	accept, err := listener.Accept()
	if err != nil {
		fmt.Println("flow: 3;启动插件出现错误:", err)
		return
	}
	encoder := gob.NewEncoder(accept)
	err = encoder.Encode(getSocketName(false))
	if err != nil {
		fmt.Println("flow: 4;启动插件出现错误:", err)
		return
	}

	// 接收数据 确定CoralBot已经成功拿到数据并返回over
	var over = false
	decoder := gob.NewDecoder(accept)
	decoder.Decode(&over)
	if !over {
		fmt.Println("flow: 5;启动插件出现错误: CoralBot服务端出现错误")
		listener.Close()
		os.RemoveAll(getSocketName(true))
		return
	}
	listener.Close()
	os.RemoveAll(getSocketName(true))
	runTask()

}

func getSocketName(t bool) string {
	// 第一次发包的名字
	if t {
		return "." + Cfg.Name + ".sock"
	}
	// 第二次发包的名字
	return fmt.Sprint("."+Cfg.Name+"@", os.Getpid(), ".sock")
}

func runTask() {
	//// 监听 Unix 域套接字
	listener, err := net.Listen("unix", getSocketName(false))
	if err != nil {
		fmt.Println("flow: 6;启动插件出现错误:", err)
		return
	}
	for {
		accept, err := listener.Accept()
		if err != nil {
			fmt.Println("flow: 7;启动插件出现错误:", err)
		}
		var bodyData []byte
		decoder := gob.NewDecoder(accept)
		decoder.Decode(&bodyData)
		switch Cfg.Type {
		case task.QQ:
			var event bot.QQEvent
			event.Explain(bodyData)
		case task.DingDing:
			var event bot.DingDingEvent
			event.Explain(bodyData)
		}
	}
}
