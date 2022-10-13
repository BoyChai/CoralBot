package main

import (
	"github.com/BoyChai/CoralBot"
)

func main() {
	var e CoralBot.Event
	CoralBot.RunCoralBot(":8080", &e)
}
