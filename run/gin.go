package run

import (
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/config"
	"github.com/BoyChai/CoralBot/plugin"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

func Run(e bot.Event, port string, readConfig bool) {
	// 创建gin对象
	g := gin.New()

	// 是否加载主配置文件
	if readConfig {
		err := config.ReadCoralBotConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// 同步配置文件配置
	if config.Cfg.Plugin {
		plugin.StartSocket()
		plugin.StartPlugin()
	}

	// 日志位置和debug日志抹除，并指定日志输出格式
	gin.DefaultWriter, gin.DebugPrintRouteFunc = logOutput(g, e)

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
		// 广播给插件
		if config.Cfg.Plugin {
			plugin.BroadcastData(bodyData)
		}
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

// 日志输出格式
func logOutput(g *gin.Engine, e bot.Event) (io.Writer, func(httpMethod, absolutePath, handlerName string, nuHandlers int)) {
	// 日志名称
	var logName = "CoralBot.log"
	// 判断日志目录是否存在
	_, err := os.Stat("logs")
	if err != nil {
		fmt.Println("日志文件夹不存再，自动创建...")
		errs := os.Mkdir("logs", 0777)
		if errs != nil {
			fmt.Println("创建日志文件夹错误: ", errs)
		}
	}

	// 打开日志文件，没有则创建(追加方式)
	logfile, _ := os.OpenFile("logs/"+logName, os.O_CREATE|os.O_APPEND, 0666)
	DefaultWriter := io.MultiWriter(logfile, os.Stdout)

	// 设置代理忽略警告
	err = g.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println("忽略代理警告错误:", err)
	}
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(e.GetLogOut(params))
		//switch e.GetType() {
		//case "DingDing":
		//	return fmt.Sprintf("[CoralBot] DingDingBot:")
		//case "QQ":
		//	return fmt.Sprintf("[CoralBot] QQBot: %s")
		//default:
		//	return fmt.Sprintf("[CoralBot]: " + e.GetType() + "类型识别上报出现错误")
		//}
	}))
	return DefaultWriter, func(httpMethod, absolutePath, handlerName string, nuHandlers int) {}
}
