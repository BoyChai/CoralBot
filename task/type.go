package task

// Task 任务载体
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
