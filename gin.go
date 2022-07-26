package CoralBot

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func RunCoralBot(port string, e *Event) {
	g := gin.Default()
	g.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		bodyData, err := ioutil.ReadAll(dataReader)
		if err != nil {
			fmt.Println(err)
		}
		e.Parse(bodyData)
	})
	g.Run(port)
}
