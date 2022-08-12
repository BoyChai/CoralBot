package main

import (
	coral "github.com/BoyChai/CoralBot"
)

func main() {
	var e coral.Event
	h := coral.Handle{
		Host:      "127.0.0.1:5700",
		Agreement: "http",
	}
	c1 := []coral.Condition{{
		Key:   &e.Message,
		Value: "hello",
		Regex: true,
	}, {
		Key:   &e.GroupID,
		Value: "你的qq群号",
	}}
	coral.NewTask(coral.Task{
		Condition: c1,
		Run: func() {
			h.SendMsg(coral.Msg{
				GroupId: e.GroupID,
				Message: "你好",
			})
		},
	})
	coral.RunCoralBot(":8080", &e)
}
