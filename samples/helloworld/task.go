package main

import "fmt"

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
