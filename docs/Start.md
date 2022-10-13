# 快速开始

## 要求

- go-cqhttp

- golang语言

## 安装

```
$ go get -u github.com/BoyChai/CoralBot
```

## 使用

使用之前需开启go-cqhttp并且指定http事件主机为本程序主机和端口，之后，创建一个名为`example.go`：

```
$ touch example.go
```

接下来，将以下代码放入`example.go`：

```
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
		Value: "<你的QQ群号>",
	}}
	coral.NewTask(coral.Task{
		Condition: c1,
		Run: func() {
			h.Reply(e, coral.Msg{
				Message: "你好",
			})
		},
	})
	coral.RunCoralBot(":8080", &e)
}

```

您可以通过以下方式运行代码`go run example.go`，运行之后可以发送hello即可收到回复。效果如下：

![效果](img/20220813102834.png)

# 基础

## 事件接收器

go-cqhttp上报的一切信息都会存在事件接收器里。创建一个事件接收器，代码如下

```go
package main

import (
	coral "github.com/BoyChai/CoralBot"
)
func main() {
    var e coral.Event
}
```

## 动作执行器

go-cqhttp的大部分行为动作都已经封装在这个对象里面，具体可参考go-cqhttp的[Api页](https://docs.go-cqhttp.org/api/#%E5%9F%BA%E7%A1%80%E4%BC%A0%E8%BE%93)。创建一个动作执行器，代码如下

```go
package main

import (
	coral "github.com/BoyChai/CoralBot"
)
func main() {
    h := coral.Handle{
        Host:      "127.0.0.1:5700",
        Agreement: "http",
    }
}
```

Host是指定go-cqhttp的ip和端口，Agreement是指定协议，目前只支持http。

## 事件触发器

当事件接收器收到上报信息就会通过事件触发器进行判断，创建一个触发器还需要了解每条上报信息的具体值，可以参考go-cqhttp的[Event页](https://docs.go-cqhttp.org/event/)。通过以下代码创建一个事件触发器

```go
package main

import (
	coral "github.com/BoyChai/CoralBot"
)
func main() {
    var e coral.Event
    c1 := []coral.Condition{{
        Key:   &e.Message,
        Value: "hello",
        Regex: true,
    }}
}
```

事件触发器c1是一个数组可以添加多个判断条件，当CoralBot运行的时候Key和Value会进行匹配，Regex是是否开启正则匹配。

## 任务的建立

```go
func NewTask(task Task) {}
```

当创建好事件触发器之后可以通过上面方法建立一个任务，建立的任务会存储在全局变量Tasks数组里面。代码如下

```go
package main

import (
	coral "github.com/BoyChai/CoralBot"
)
func main() {
	var e coral.Event
	c1 := []coral.Condition{{
		Key:   &e.Message,
		Value: "hello",
		Regex: true,
	}}
	coral.NewTask(coral.Task{
		Condition: c1,
		Run: func() {
			fmt.Println("hello事件触发了")
		},
	})
}
```

## 运行CoralBot

```go
func RunCoralBot(port string, e *Event) {}
```

通过上面函数来运行CoralBot。需要传入一个端口号和事件监听器，运行时CoralBot会把go-cqhttp的全部事件都王事件监听器里面存放。

```go
package main

import (
	coral "github.com/BoyChai/CoralBot"
)
func main() {
	var e coral.Event
	c1 := []coral.Condition{{
		Key:   &e.Message,
		Value: "hello",
		Regex: true,
	}}
	coral.NewTask(coral.Task{
		Condition: c1,
		Run: func() {
			fmt.Println("hello事件触发了")
		},
	})
	coral.RunCoralBot(":8080", &e)
}
```

#  其他

## 概述

CoralBot对于go-cqhttp的APi基本全部封装到了Handle对象里面了，具体的变更内容可以看下面的[API的封装](#API的封装)，对于go-cqhttp的上报事件信息封装到了Event这个结构体中，具体的更不内容可以看下面的[Event的封装](#Event的封装)。

对于CoralBot的文档还有待完善。

## API的封装

CoralBot的对于go-cqhttp的API封装的名称全部变成了一下格式

```
//go-cqhttp
/send_private_msg

//CoralBot
func (h Handle) SendPrivateMsg(m Msg) (map[string]interface{}, error) { }
```

## Event的封装

CoralBot的对于go-cqhttp的Event返回的json数据全部封装到了Event这个结构体里面，例如下面

```
//go-cqhttp
jsonStr{
time
self_id
post_type
}

//CoralBot
Event.Time
Event.SelfID
Event.PostType
```

## Handle对象

Handle的每一个对api封装的函数都会返回一个json对象和一个error，如果请求成功会返回一个json，json的内容可以参考[go-cqhttp的api文档](https://docs.go-cqhttp.org/api/)

## 频道相关

频道上报的信息，user_id和message_id和普通消息存在类型冲突，所以频道的这两个值的封装如下

```
Event.GuildUserID
Event.GuildMessageID
```