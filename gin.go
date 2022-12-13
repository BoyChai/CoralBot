package CoralBot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func RunCoralBot(port string, e *Event) {
	// 空事件
	var init Event
	// 日志时间
	now := time.Now()
	// 创建gin对象
	g := gin.New()
	// 日志名称组装
	logName := fmt.Sprintf("logs/" + "CoralBot" + now.Format("20060102-150405") + ".log")
	errs := os.Mkdir("logs", 0777)
	if errs != nil {
		fmt.Println(errs)
	}
	logfile, _ := os.Create(logName)
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[CoralBot] QQBot:%d 时间:%s 上报类型:%s 事件内容为:%+v\n",
			e.SelfID,
			params.TimeStamp.Format(time.RFC3339),
			e.PostType,
			e,
		)
	}))
	//加载主配置文件
	err := readCoralBotConfig()
	if err != nil {
		return
	}
	//同步配置文件配置
	if Cfg.Plugin {
		// 加载插件
		err = e.loadPlugin()
		if err != nil {
			fmt.Println("插件加载失败：", err)
		}
	}
	// 接收上报
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := ioutil.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		//e.Parse(bodyData)
		*e = init
		//e.bodyData = bodyData
		e.explain(bodyData)
	})
	err = g.Run(port)
	if err != nil {
		fmt.Printf("gin:%v", err)
	}
}
