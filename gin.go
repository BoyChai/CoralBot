package CoralBot

import (
	"fmt"
	"github.com/BoyChai/CoralBot/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func RunCoralBot(port string, e *Event, readConfig bool) {
	// 空事件
	var init Event
	// 创建gin对象
	g := gin.New()
	// 是否加载主配置文件
	if readConfig {
		err := readCoralBotConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//同步配置文件配置
	if Cfg.Plugin {
		// 加载插件
		err := e.loadPlugin()
		if err != nil {
			fmt.Println("插件加载失败：", err)
		}
	}

	// 日志位置和debug日志抹除，并指定日志输出格式
	gin.DefaultWriter, gin.DebugPrintRouteFunc = utils.LogOutput(g, e)

	// 接收上报
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := ioutil.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		*e = init
		e.explain(bodyData)
	})
	// 设置代理忽略警告
	err := g.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 选中端口
	if port == "" {
		err = g.Run(fmt.Sprint(":", Cfg.Listen))
	} else {
		err = g.Run(port)
	}
	if err != nil {
		fmt.Printf("gin:%v", err)
	}
}
