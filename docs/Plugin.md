# 插件相关

## 编写插件

```go
package main

import (
	"github.com/BoyChai/CoralBot"
	"github.com/dullgiulio/pingo"
)

var H CoralBot.Handle

var Info CoralBot.PluginInfo

type MyPlugin struct{}

func (p *MyPlugin) PluginInfo(n string, info *CoralBot.PluginInfo) error {
	*info = Info
	return nil
}

func (p *MyPlugin) GetPlugin(e *CoralBot.Event, Task *[]CoralBot.Task) error {
	CoralBot.NewPluginTask(CoralBot.Task{
		Condition: []CoralBot.Condition{
			{
				Key:   &e.Message,
				Value: "插件测试",
			},
		},
		RunName: "MyPlugin.Run1",
	})
	*Task = CoralBot.Tasks
	return nil
}

func (p *MyPlugin) Run1(e CoralBot.Event, n *string) error {
	h := CoralBot.Handle{
		Host:      "127.0.0.1:5700",
		Agreement: "http",
	}
	h.Reply(e, CoralBot.Msg{
		Message: "插件激活成功",
	})
	return nil
}
func main() {
	H.Host = "127.0.0.1:5700"
	H.Agreement = "http"
	Info = CoralBot.PluginInfo{
		Name:      "test_demo",
		Summary:   "测试插件demo",
		Version:   "v0.0.1",
		Developer: "BoyChai",
		Email:     "1972567225@qq.com",
	}

	plugin := &MyPlugin{}

	pingo.Register(plugin)
	pingo.Run()
}
```

编写需注意：

- 插件结构体必须叫做"MyPlugin"
- 必须拥有以下方法
   - PluginInfo：用来获取插件信息
  - GetPlugin：用来获取插件任务
- Task.RunName的值应该为"MyPlugin.方法"
- 当任务触发成功之后就会执行Task.RunName指定的方法
- 编译之后后缀需要改为coral
- 插件的动作执行器(Handle)需要自己创建,目前不可以使用主程序的

## 使用插件

把编译好的插件放入"./plugin"文件夹,之后代码如下

```go
package main

import (
	"github.com/BoyChai/CoralBot"
)

func main() {
	var e CoralBot.Event
	CoralBot.RunCoralBot(":8080", &e)
}
```

运行之后会进行加载插件,并输出插件信息

```bash
2022/10/13 17:43:16 tcp
===============已加载插件===============
插件名称： test_demo
插件版本： v0.0.1
插件概述： 测试插件demo
插件作者： BoyChai
作者邮箱： 1972567225@qq.com
```

运行之后执行，发送消息"插件测试"即可，效果如下：

![效果](./img/20221013173153.png)