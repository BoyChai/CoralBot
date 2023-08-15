package main

import (
	"github.com/BoyChai/CoralBot/bot"
	"github.com/BoyChai/CoralBot/run"
)

func main() {
	event := bot.QQEvent{}
	run.Run(&event, ":8080", true)
}
