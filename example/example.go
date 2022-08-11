package main

import (
	"fmt"
	coral "github.com/BoyChai/CoralBot"
	"strconv"
)

func main() {
	var e coral.Event
	h := coral.Handle{
		Host: "127.0.0.1:5700",
	}
	c1 := []coral.Condition{{
		Key:   &e.Message,
		Value: "hello",
		Regex: true,
	}, {
		Key:   &e.GroupID,
		Value: "群号",
	}}
	coral.NewTask(coral.Task{
		Mode:      "all_message",
		Condition: c1,
		Run: func() {
			groupId, err := strconv.ParseInt(e.GroupID, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			h.SendMsg(coral.Msg{
				GroupId: groupId,
				Message: "你好",
			})
		},
	})
	coral.RunCoralBot(":8080", &e)
}
