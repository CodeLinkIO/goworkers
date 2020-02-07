package notifiers

import (
	"github.com/CodeLinkIO/goworkers/tasks"
	"github.com/CodeLinkIO/goworkers/triggers"
)

type notifier struct {
	tasks  chan tasks.Task
	errors chan error
}

func NewTrigger() triggers.Trigger {
	return &notifier{
		tasks:  make(chan tasks.Task),
		errors: make(chan error),
	}
}

func (t *notifier) Listen() chan tasks.Task {
	return t.tasks
}

func (t *notifier) NotifyError() chan error {
	return t.errors
}

func (t *notifier) Notify(task tasks.Task) {
	t.tasks <- task
}
