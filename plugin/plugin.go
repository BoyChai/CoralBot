package plugin

import (
	"fmt"
	"github.com/BoyChai/CoralBot/task"
	"net"
	"os"
)

var Infos []task.Info

var accepts []net.Conn

func StartSocket() {
	socketFile := "CoralBot.sock"

	// 删除已经存在的套接字文件
	if err := os.RemoveAll(socketFile); err != nil {
		fmt.Println("removing plugin socket file:", err)
		return
	}

	// 监听 Unix 域套接字
	listener, err := net.Listen("unix", socketFile)
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
		Infos = append(Infos, ReadInfo(conn))
	}
}

func BroadcastData(data []byte) {
	for _, accept := range accepts {
		WriteData(data, accept)
	}
}
