package CoralBot

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func RunCoralBot(port string, e *Event) {
	var init Event
	g := gin.Default()
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := ioutil.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		//e.Parse(bodyData)
		*e = init
		e.bodyData = string(bodyData)
		e.explain()
	})
	err := g.Run(port)
	if err != nil {
		fmt.Printf("gin:%v", err)
	}
}
