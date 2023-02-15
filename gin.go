package CoralBot

import (
	"fmt"
	"github.com/BoyChai/CoralBot/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func RunCoralBot(port string, e *Event, readConfig bool) {
	// 重置事件的空事件
	var init Event

	// 创建gin对象
	g := gin.New()

	// 是否加载主配置文件
	if readConfig {
		err := utils.ReadCoralBotConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//同步配置文件配置
	if utils.Cfg.Plugin {
		// 加载插件
		err := e.loadPlugin()
		if err != nil {
			fmt.Println("插件加载失败：", err)
		}
	}

	// 日志位置和debug日志抹除，并指定日志输出格式
	gin.DefaultWriter, gin.DebugPrintRouteFunc = logOutput(g, e)

	// 接收上报
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := io.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		*e = init
		e.explain(bodyData)
	})

	// 选择端口并启动程序
	var err error
	if port == "" {
		err = g.Run(fmt.Sprint(":", utils.Cfg.Listen))
	} else {
		err = g.Run(port)
	}
	if err != nil {
		fmt.Printf("gin:%v", err)
	}
}

func logOutput(g *gin.Engine, e *Event) (io.Writer, func(httpMethod, absolutePath, handlerName string, nuHandlers int)) {
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

	// 日志格式
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[CoralBot] QQBot:%d 时间:%s 上报类型:%s 事件内容为:%+v\n",
			e.SelfID,
			params.TimeStamp.Format(time.RFC3339),
			e.PostType,
			e,
		)
	}))

	return DefaultWriter, func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	}
}
