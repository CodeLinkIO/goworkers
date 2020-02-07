package main

import "fmt"

type sampleTask struct {
	Message string
}

func (t *sampleTask) Type() string {
	return "Hello"
}

func (t *sampleTask) Complete() {
	fmt.Println(fmt.Sprintf("Task %v complete", t.Message))
}

func (t *sampleTask) Fail(err error) {
	fmt.Println(fmt.Sprintf("Task %v fail: %v", t.Message, err.Error()))
}
