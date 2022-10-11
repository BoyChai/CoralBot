package CoralBot

import (
	"fmt"
	"github.com/dullgiulio/pingo"
)

// Task 创建一个任务
type Task struct {
	Condition   []Condition
	Run         func()
	plugin      bool
	pingoServer *pingo.Plugin
	RunName     string
	info        PluginInfo
}

// Condition 触发条件
type Condition struct {
	Key   interface{}
	Value string
	Regex bool
}

var Tasks []Task

// NewTask 创建一个动作
func NewTask(task Task) {
	task.plugin = false
	Tasks = append(Tasks, task)
}

// NewPluginTask 创建一个插件动作
func NewPluginTask(task Task) {
	task.plugin = true
	if task.RunName == "" {
		fmt.Println("创建插件动作错误：RunName值为空")
		return
	}
	Tasks = append(Tasks, task)
}
