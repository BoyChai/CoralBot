package bot

type Event interface {
	GetType() string
	GetDocs() string
	GetRunName() string
	SetRunName(string)
	Explain(bodyData []byte)
}

// Other 其他
type Other struct {
	// 运行插件
	RunName string
}
