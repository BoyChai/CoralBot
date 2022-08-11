package CoralBot

// Task 创建一个任务
type Task struct {
	//触发条件类型
	Mode string
	//Message string
	//触发条件
	Condition []Condition
	Run       func()
}

// Condition 触发条件
type Condition struct {
	Key   *string
	Value string
	Regex bool
}

var Tasks []Task

// NewTask 创建一个动作
func NewTask(task Task) {
	Tasks = append(Tasks, task)
}
