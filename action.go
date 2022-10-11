package CoralBot

import "fmt"

// Task 创建一个任务
type Task struct {
	Condition []Condition
	Run       func()
	Plugin    bool
	RunName   string
}

// Condition 触发条件
type Condition struct {
	Key   interface{}
	Value string
	Regex bool
}

// Plugin 插件信息
//type Plugin struct {
//	// 插件名称
//	Name string
//	// 插件简介
//	Summary string
//	// 插件版本
//	Version string
//	// 插件开发者
//	Developer string
//	// 插件开发者邮箱
//	Email string
//	// 执行调用函数
//	RunName string
//}

var Tasks []Task

// NewTask 创建一个动作
func NewTask(task Task) {
	task.Plugin = false
	Tasks = append(Tasks, task)
}

// NewPluginTask 创建一个插件动作
func NewPluginTask(task Task) {
	task.Plugin = true
	if task.RunName == "" {
		fmt.Println("创建插件动作错误：RunName值为空")
		return
	}
	Tasks = append(Tasks, task)
}
