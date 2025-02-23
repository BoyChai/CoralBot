package botv1

type Evnet interface {
	GetType() string
	Parse(jsonStr string)
	Broadcast() error
}
