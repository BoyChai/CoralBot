package utils

import (
	"fmt"
	"github.com/BoyChai/CoralBot"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

var logName = "CoralBot.log"

func LogOutput(g *gin.Engine, e *CoralBot.Event) (io.Writer, func(httpMethod, absolutePath, handlerName string, nuHandlers int)) {
	// 判断日志目录是否存在
	_, err := os.Stat("logs")
	if err == nil {
		fmt.Println("日志文件夹不存再，自动创建...")
		errs := os.Mkdir("logs", 0777)
		if errs != nil {
			fmt.Println("创建日志文件夹错误: ", errs)
		}
	}

	// 打开日志文件，没有则创建(追加方式)
	logfile, _ := os.OpenFile("logs/"+logName, os.O_CREATE|os.O_APPEND, 0666)
	DefaultWriter := io.MultiWriter(logfile, os.Stdout)

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
