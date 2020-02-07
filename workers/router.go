package workers

import (
	"context"
	"goworkers/tasks"
)

type Handler func(ctx context.Context, task tasks.Task) error

type Router interface {
	Register(taskType string, handler Handler)
	Resolve(taskType string) (Handler, bool)
}

func NewRouter() Router {
	return make(router)
}

type router map[string]Handler

func (r router) Register(taskType string, handler Handler) {
	r[taskType] = handler
}

func (r router) Resolve(taskType string) (Handler, bool) {
	handler := r[taskType]
	return handler, handler != nil
}
