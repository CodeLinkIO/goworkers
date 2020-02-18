# goworkers
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

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
## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/huynvk"><img src="https://avatars2.githubusercontent.com/u/15973503?v=4" width="100px;" alt=""/><br /><sub><b>Huy Ngo</b></sub></a><br /><a href="#ideas-huynvk" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://github.com/peterphan1996"><img src="https://avatars1.githubusercontent.com/u/28189578?v=4" width="100px;" alt=""/><br /><sub><b>Peter Phan</b></sub></a><br /><a href="#ideas-peterphan1996" title="Ideas, Planning, & Feedback">🤔</a> <a href="https://github.com/CodeLinkIO/goworkers/commits?author=peterphan1996" title="Documentation">📖</a> <a href="https://github.com/CodeLinkIO/goworkers/commits?author=peterphan1996" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!