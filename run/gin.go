package run

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/config"
	"github.com/BoyChai/CoralBot/log"
	"github.com/BoyChai/CoralBot/plugin"
)

func Run(e bot.Event, port string, readConfig bool) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 检测请求方式
		if r.Method != "POST" {
			log.Debug("Method not allowed: " + r.Method)
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
		// 释放资源
		defer r.Body.Close()
		// 数据过滤
		if e.GetType() == "DingDing" {
			timestamp, err := strconv.ParseInt(r.Header.Get("timestamp"), 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			config.Timestamp = timestamp
			config.Sign = r.Header.Get("sign")
			// 如果需要在函数结束时释放资源，应该使用defer语句
			defer r.Header.Clone()
		}
		log.Debug(string(body))
		e.Explain(body)
		// 广播给插件
		if config.Cfg.Plugin {
			plugin.BroadcastData(body)
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
		log.Fatal("HttpServer:", err.Error)
		return
	}
}
