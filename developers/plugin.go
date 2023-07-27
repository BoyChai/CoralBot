package developers

import (
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/plugin"
	"net"
)

func RunPlugin(event bot.Event) {
	SocketFile := "./CoralBot.sock"
	//conn, err := net.Dial("unix", config.SocketFile)
	conn, err := net.Dial("unix", SocketFile)
	if err != nil {
		fmt.Println("socket连接失败。请检查CoralBot是否启动。")
		return
	}
	plugin.WriteInfo(information, conn)

	// 等待播报
	for {
		data := plugin.ReadData(conn)
		if string(data) == "Heartbeat" {
			continue
		}
		event.Explain(data)
	}
}
