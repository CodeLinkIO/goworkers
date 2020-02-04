package triggers

import "goworkers/tasks"

type Trigger interface {
	Notify(tasks.Task)
	Listen() chan tasks.Task
	NotifyError() chan error
}
