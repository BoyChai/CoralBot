# CoralBot

## 概述

基于基于go-cqhttp的后端开发库

## 快速开始

### 要求

- go 1.18以上

### 安装

```
$ go get -u github.com/BoyChai/CoralBot
```

### 使用

使用之前需开启go-cqhttp并且指定http事件主机为本程序主机和端口，之后，创建一个名为`example.go`：

```
$ touch example.go
```

接下来，将以下代码放入`example.go`：

```
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
```

您可以通过以下方式运行代码`go run example.go`，运行之后可以在群聊里面发送hello即可收到回复。