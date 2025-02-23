package receiver

import (
	"fmt"
	"strings"
)

var colors = []string{
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
}

var textLines = []string{
	"_________                   _____________      _____ ",
	"__  ____/__________________ ___  /__  __ )_______  /_",
	"_  /    _  __ \\_  ___/  __ `/_  /__  __  |  __ \\  __/",
	"/ /___  / /_/ /  /   / /_/ /_  / _  /_/ // /_/ / /_  ",
	"\\____/  \\____//_/    \\__,_/ /_/  /_____/ \\____/\\__/  ",
	"",
}

// 生成彩色文本
func coloredText() string {
	var coloredLines []string
	for i, line := range textLines {
		if i < len(colors) { // 确保不越界
			coloredLines = append(coloredLines, colors[i]+line+"\033[0m")
		} else {
			coloredLines = append(coloredLines, line) // 空行保持无色
		}
	}
	return strings.Join(coloredLines, "\n")
}

func init() {
	fmt.Println(coloredText())
}
