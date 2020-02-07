# goworkers

goworkers is a library in Golang that allow users launch and manage workers


## Installation

```go
go get github.com/CodelinkIO/goworkers
```

## Example
You can find more samples with different complexities inside `samples` folder. Here is a simple sample of how to use:

### Define a task

```go
type helloTask struct {
	Message string
}

func (t *helloTask) Type() string {
	return "Hello"
}

func (t *helloTask) Complete() {
	fmt.Println(fmt.Sprintf("Task %v complete", t.Message))
}

func (t *helloTask) Fail(err error) {
	fmt.Println(fmt.Sprintf("Task %v fail: %v", t.Message, err.Error()))
}
```

### Define a handler

```go
func handleHello(ctx context.Context, task tasks.Task) error {
	workerContext, ok := ctx.Value(workers.WorkerContextKey).(workers.WorkerContext)
	if !ok {
		return fmt.Errorf("Cannot get worker context")
	}
	processingTask, ok := task.(*helloTask)
	if !ok {
		return fmt.Errorf("Cannot parse task")
	}

	fmt.Println(fmt.Sprintf("%v: Handle task %v %v", workerContext.ID, processingTask.Type(), processingTask.Message))
	time.Sleep(200 * time.Millisecond)
	return nil
}
```

### Setup controller

```go
options := workers.ControllerOptions{
    NumOfWorker: 10,
}
ctx := context.Background()

router := workers.NewRouter()
router.Register("Hello", handleHello)

trigger := notifiers.NewTrigger()

controller := workers.NewController(ctx, trigger, router, options)
controller.Run()
```

### Trigger the task from another goroutine

```go
for i := 0; i < 20; i++ {
    trigger.Notify(&helloTask{Message: fmt.Sprintf("tsk_%d", i)})
    time.Sleep(100 * time.Millisecond)
}
```