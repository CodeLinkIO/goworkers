package workers

import (
	"context"
	"goworkers/tasks"
)

type Handler func(ctx context.Context, task tasks.Task) error

type Router map[string]Handler

func (r *Router) Register(taskType string, handler Handler) {
	(*r)[taskType] = handler
}

func (r *Router) Resolve(taskType string) (Handler, bool) {
	handler := (*r)[taskType]
	return handler, handler != nil
}
