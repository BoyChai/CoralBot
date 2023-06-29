package task

var Tasks []Task

// Task 任务载体
type Task struct {
	Condition []Condition
	Run       func()
	//Info      Info
	//Plugin    bool
}

// Condition 触发条件
type Condition struct {
	Key   interface{}
	Value string
	Regex bool
}

// NewTask 创建任务
func NewTask(task Task) {
	Tasks = append(Tasks, task)
}

// NewPluginTask 创建插件任务
//func NewPluginTask(task Task) {
//	task.Plugin = true
//	Tasks = append(Tasks, task)
//}
