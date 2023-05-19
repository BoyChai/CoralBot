package developers

import (
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/plugin"
	"net"
)

//func (p *Plugin) PluginInfo(n string, info *task.Info) error {
//	*info = information
//	c, err := readConfig()
//	if err != nil {
//		return err
//	}
//	switch info.Type {
//	case task.QQ:
//		h := GetQQHandler()
//		h.Host = c.Host
//		h.Agreement = c.Agreement
//	case task.DingDing:
//		h := GetDingDingHandler()
//		h.AppKey = c.AppKey
//		h.AppSecret = c.AppSecret
//	}
//
//	return nil
//}

//func (p *Plugin) GetPlugin(e interface{}, Task *[]task.Task) error {
//	//func (p *Plugin) GetPlugin(e *bot.Event, Task *[]task.Task) error {
//	event = e
//	*Task = task.Tasks
//	return nil
//}

//func (p *Plugin) Handles(e interface{}, n *string) error {
//	//func (p *Plugin) Handles(e *bot.QQEvent, n *string) error {
//	//func (p *Plugin) Handles(e bot.Event, n *string) error {
//	for _, value := range handles {
//		if e.(*bot.QQEvent).GetRunName() == value.name {
//			err := value.run(e)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}

func RunPlugin(event bot.Event) {
	SocketFile := "./CoralBot.sock"
	//conn, err := net.Dial("unix", config.SocketFile)
	conn, err := net.Dial("unix", SocketFile)
	if err != nil {
		fmt.Println("socket连接失败。请检查CoralBot是否启动。")
		return
	}
	plugin.WriteInfo(information, conn)
	data := plugin.ReadData(conn)
	event.Explain(data)
}
