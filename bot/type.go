package bot

import "github.com/gin-gonic/gin"

type Event interface {
	GetType() string
	GetDocs() string
	GetRunName() string
	SetRunName(string)
	Explain(bodyData []byte)
	GetLogOut(params gin.LogFormatterParams) string
}

// Other 其他
type Other struct {
	// 运行插件
	RunName string
}
