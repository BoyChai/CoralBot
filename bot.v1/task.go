package botv1

var Tasks []func(Evnet) error

// NewTask 创建任务
func NewTask(task func(Evnet) error) {
	Tasks = append(Tasks, task)
}
