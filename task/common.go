package task

var Tasks []Task

// NewTask 创建任务
func NewTask(task Task) {
	Tasks = append(Tasks, task)
}
