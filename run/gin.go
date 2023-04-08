package run

import (
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/config"
	"github.com/BoyChai/CoralBot/task"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

func Run(e bot.Event, port string, readConfig bool) {
	// 创建gin对象
	g := gin.Default()

	// 是否加载主配置文件
	if readConfig {
		err := config.ReadCoralBotConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//同步配置文件配置
	if config.Cfg.Plugin {
		// 加载插件
		err := task.LoadPlugin()
		if err != nil {
			fmt.Println("插件加载失败：", err)
		}
	}

	// 接收上报
	g.POST("/", func(c *gin.Context) {
		var err error
		dataReader := c.Request.Body
		bodyData, err := io.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
			err = nil
		}
		if e.GetType() == "DingDing" {
			timestamp, err := strconv.ParseInt(c.Request.Header.Get("timestamp"), 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			config.Timestamp = timestamp
			config.Sign = c.Request.Header.Get("sign")
			defer c.Request.Header.Clone()
		}
		e.Explain(bodyData)
	})
	// 选择端口并启动程序
	var err error
	if port == "" {
		err = g.Run(fmt.Sprint(":", config.Cfg.Listen))
	} else {
		err = g.Run(port)
	}
	if err != nil {
		fmt.Printf("gin:%v", err)
	}
}
