package CoralBot

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 主页面
func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{"item": Tasks})
}

// Control 渲染控制
func Control(e *gin.Engine) {
	// 映射资源
	e.Static("/static/", "template/static")
	// 导入模板
	e.LoadHTMLGlob("template/*.tmpl")
	// 主界面
	e.GET("/", home)
}
