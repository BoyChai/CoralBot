package bot

type Event interface {
	GetType() string
	GetDocs() string
	Explain(bodyData []byte)
}
