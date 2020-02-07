package triggers

import "github.com/CodeLinkIO/goworkers/tasks"

type Trigger interface {
	Notify(tasks.Task)
	Listen() chan tasks.Task
	NotifyError() chan error
}
