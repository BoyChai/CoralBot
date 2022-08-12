package CoralBot

// Task 创建一个任务
type Task struct {
	Condition []Condition
	Run       func()
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
	Tasks = append(Tasks, task)
}
