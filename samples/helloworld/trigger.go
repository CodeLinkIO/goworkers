package main

import "goworkers/tasks"

type trigger struct {
	tasks  chan tasks.Task
	errors chan error
}

func NewTrigger() *trigger {
	return &trigger{
		tasks:  make(chan tasks.Task),
		errors: make(chan error),
	}
}

func (t *trigger) Listen() chan tasks.Task {
	return t.tasks
}

func (t *trigger) NotifyError() chan error {
	return t.errors
}

func (t *trigger) Notify(task tasks.Task) {
	t.tasks <- task
}
