package main

import "github.com/BoyChai/CoralBot"

func main() {
	var e CoralBot.Event
	h := CoralBot.Handle{
		Host: "127.0.0.1:5700",
	}
	c1 := []CoralBot.Condition{{
		Key:   &e.Message,
		Value: "hello",
		Regex: true,
	}, {
		Key:   &e.Message,
		Value: "你好",
	}}
	CoralBot.NewTask(CoralBot.Task{
		Mode:      "all_message",
		Condition: c1,
		Run: func() {
			h.SendMsg("", e.GroupID, "你好", "")
		},
	})
	CoralBot.RunCoralBot(":8080", &e)
}
