package receiver

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	botv1 "github.com/BoyChai/CoralBot/bot.v1"
	"github.com/BoyChai/CoralBot/config"
	"github.com/BoyChai/CoralBot/log"
	"github.com/BoyChai/CoralBot/plugin"
)

func StartHttpServer(e botv1.Evnet, port string, readConfig bool) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 检测请求方式
		if r.Method != "POST" {
			log.Debug("Method not allowed: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// 数据获取
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("Failed to read request body")
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		log.Debug("%s", string(body))
		defer r.Body.Close()

		// 数据组装
		if e.GetType() == "DingDing" {
			timestamp, err := strconv.ParseInt(r.Header.Get("timestamp"), 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			config.Timestamp = timestamp
			config.Sign = r.Header.Get("sign")
			defer r.Header.Clone()
		}
		// 解析器
		e.Parse(string(body))
		// 广播
		err = e.Broadcast()
		if err != nil {
			log.Error("%s", err.Error())
			return
		}
	})

	// 是否加载主配置文件
	if readConfig {
		err := config.ReadCoralBotConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// 同步配置文件配置
	if config.Cfg.Plugin {
		plugin.StartSocket()
		plugin.StartPlugin()
	}

	// 选择端口并启动程序
	var err error
	if port == "" {
		err = http.ListenAndServe(fmt.Sprint(":", config.Cfg.Listen), nil)

	} else {
		err = http.ListenAndServe(port, nil)
	}
	if err != nil {
		log.Fatal("HttpServer: %s", err.Error())
		return
	}
}
