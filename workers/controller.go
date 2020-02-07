package workers

import (
	"context"
	"fmt"
	"goworkers/tasks"
	"goworkers/triggers"
	"goworkers/utils/errors"

	"golang.org/x/sync/errgroup"
)

type ControllerOptions struct {
	NumOfWorker int
}

type Controller interface {
	Run() error
	Stop()
}

func NewController(ctx context.Context, trigger triggers.Trigger, router Router,
	options ControllerOptions) Controller {
	return &controller{
		router:    router,
		options:   options,
		trigger:   trigger,
		parentCtx: ctx,
	}
}

type controller struct {
	router    Router
	trigger   triggers.Trigger
	options   ControllerOptions
	parentCtx context.Context
	exCtx     context.Context
	cancel    context.CancelFunc
	isRunning bool
	eg        errgroup.Group
	taskChan  chan tasks.Task
}

func (c *controller) Run() error {
	if c.isRunning {
		return nil
	}

	c.exCtx, c.cancel = context.WithCancel(c.parentCtx)
	c.isRunning = true
	defer c.Stop()

	errorChan := c.trigger.NotifyError()
	go func() {
		for {
			if _, ok := <-errorChan; ok {
				c.Stop()
			}
		}
	}()
	c.taskChan = c.trigger.Listen()

	for i := 0; i < c.options.NumOfWorker; i++ {
		workerCtx := context.WithValue(c.exCtx, WorkerContextKey, WorkerContext{
			ID: fmt.Sprintf("Worker %d", i),
		})
		c.eg.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = errors.ConvertPanicToError(r)
				}
			}()

			for {
				select {
				case <-workerCtx.Done():
					return nil

				case task := <-c.taskChan:
					handler, ok := c.router.Resolve(task.Type())
					if !ok {
						task.Fail(fmt.Errorf("Cannot handle task type %s", task.Type()))
					}

					err := handler(workerCtx, task)
					if err != nil {
						task.Fail(err)
					}
					task.Complete()
				}

			}
		})
	}

	return c.eg.Wait()
}

func (c *controller) Stop() {
	c.cancel()
	c.isRunning = false
}
