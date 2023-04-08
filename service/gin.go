package service

import (
	"fmt"
	"github.com/BoyChai/CoralBot/bot"
	"github.com/gin-gonic/gin"
	"io"
)

func Run(e bot.Event, port string) {
	// 创建gin对象
	g := gin.Default()

	// 接收上报
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := io.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		e.Explain(bodyData)
	})
	g.Run(port)
}
