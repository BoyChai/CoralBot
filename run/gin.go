package run

import (
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/config"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

func Run(e bot.Event, port string) {
	// 创建gin对象
	g := gin.Default()

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
	g.Run(port)
}
