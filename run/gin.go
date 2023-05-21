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
	g := gin.Default()

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
		// 加载插件
		//err := task.LoadPlugin()
		//if err != nil {
		//	fmt.Println("插件加载失败：", err)
		//}
		plugin.StartSocket()
		plugin.StartPlugin()
	}

	// 日志位置和debug日志抹除，并指定日志输出格式
	gin.DefaultWriter, gin.DebugPrintRouteFunc = logOutput(g, e.GetType())
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[CoralBot] %s | %s | %s | %s ",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))
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
func logOutput(g *gin.Engine, bodyType string) (io.Writer, func(httpMethod, absolutePath, handlerName string, nuHandlers int)) {
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

	g.Use(func(context *gin.Context) {
		switch bodyType {
		case "DingDing":
			fmt.Println("[CoralBot] DingDingBot:")
		case "QQ":
			fmt.Println("[CoralBot] QQBot:")
		default:
			fmt.Println("[CoralBot]: " + bodyType + "类型识别上报出现错误")
		}
	})

	// 日志格式
	//g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	switch e.(type) {
	//	case *bot.QQEvent:
	//		return fmt.Sprintf("[CoralBot] QQBot:%d 时间:%s 上报类型:%s 事件内容为:%+v\n",
	//			e.(*bot.QQEvent).SelfID,
	//			params.TimeStamp.Format(time.RFC3339),
	//			e.(*bot.QQEvent).PostType,
	//			//e.(*bot.QQEvent).Message,
	//			e.(*bot.QQEvent),
	//		)
	//	case *bot.DingDingEvent:
	//		return fmt.Sprintf("[CoralBot] DingDingBot:%d 时间:%s 上报类型:%s 事件内容为:%+v\n",
	//			e.(*bot.DingDingEvent).ChatbotCorpId,
	//			params.TimeStamp.Format(time.RFC3339),
	//			e.(*bot.DingDingEvent).Msgtype,
	//		)
	//	default:
	//		return fmt.Sprintf("[CoralBot] Log: 未知上报类型")
	//	}
	//
	//}))

	return DefaultWriter, func(httpMethod, absolutePath, handlerName string, nuHandlers int) {}
}
