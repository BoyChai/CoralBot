package plugin

import (
	"fmt"
	"github.com/BoyChai/CoralBot/config"
	"github.com/BoyChai/CoralBot/task"
	"net"
	"os"
)

var Infos []task.Info

var accepts []net.Conn

func StartSocket() {
	SocketFile := "CoralBot.sock"
	// 删除已经存在的套接字文件
	//if err := os.RemoveAll(config.SocketFile); err != nil {
	if err := os.RemoveAll(SocketFile); err != nil {
		fmt.Println("removing plugin socket file:", err)
		return
	}

	// 监听 Unix 域套接字
	listener, err := net.Listen("unix", SocketFile)
	if err != nil {
		fmt.Println("Error listening on plugin socket:", err)
		return
	}
	go receiveInformation(listener)
}

func receiveInformation(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		accepts = append(accepts, conn)
		if err != nil {
			fmt.Println("Error accepting plugin connection:", err)
			continue
		}
		info := ReadInfo(conn)
		Infos = append(Infos, info)
		if config.Cfg.PluginInfo {
			fmt.Println("Loading succeeded：", info.Name)
			fmt.Println("===============Plugin-Info===============")
			fmt.Println("插件名称：", info.Name)
			fmt.Println("插件版本：", info.Version)
			fmt.Println("插件概述：", info.Summary)
			fmt.Println("插件作者：", info.Developer)
			fmt.Println("作者邮箱：", info.Email)
			fmt.Println("=========================================")
			fmt.Println("CoralBot加载插件数量为：", len(Infos))
		} else {
			fmt.Println("Loading succeeded:", info.Name)
			fmt.Println("CoralBot加载插件数量为：", len(Infos))
		}

	}
}

func BroadcastData(data []byte) {
	for _, accept := range accepts {
		WriteData(data, accept)
	}
}
