package CoralBot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func RunCoralBot(port string, e *Event) {
	var init Event
	g := gin.New()
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		bodyData, err := ioutil.ReadAll(params.Request.Body)
		if err != nil {
			fmt.Println("logs:", err)
		}
		err = params.Request.Body.Close()
		if err != nil {
			fmt.Println("logs:", err)
		}
		var data Event
		err = json.Unmarshal(bodyData, &data)
		if err != nil {
			fmt.Println("logs:", err)
		}
		return fmt.Sprintf("[CoralBot] QQBot:%s 时间:%s 上报类型:%s",
			data.SelfID,
			params.TimeStamp.Format(time.RFC3339),
			data.PostType,
		)
	}))
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := ioutil.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		//e.Parse(bodyData)
		*e = init
		e.bodyData = bodyData
		e.explain()
	})
	err := g.Run(port)
	if err != nil {
		fmt.Printf("gin:%v", err)
	}
}
