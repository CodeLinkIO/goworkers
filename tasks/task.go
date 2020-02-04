package tasks

type Task interface {
	Type() string
	Complete()
	Fail(error)
}
