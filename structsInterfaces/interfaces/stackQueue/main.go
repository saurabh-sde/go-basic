package main

import (
	"fmt"

	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/stackQueue/queue"
	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/stackQueue/stack"
)

type DS interface {
	Pop() int
	Push(n int)
	Print()
}

func NewDS(dataType string) DS {
	fmt.Println("Implementing DS as", dataType)
	if dataType == "stack" {
		return &stack.Stack{}
	} else if dataType == "queue" {
		return &queue.Queue{}
	}
	fmt.Println("invalid type", dataType)
	return nil
}

func Push(p DS, n int) {
	p.Push(n)
}

func Pop(p DS) {
	fmt.Println("Pop: ", p.Pop())
}
func Print(p DS) {
	p.Print()
}

//implement an integer stack which satisfies the popper interface.
//if stack is empty return -1.
//you can only use PopThree function to pop the elements.
//you cannot modify the interface and function.

func main() {
	// Using Stack struct
	fmt.Println("Stack")
	newStack := &stack.Stack{}
	newStack.Push(1)
	newStack.Push(2)
	newStack.Push(3)
	Push(newStack, 4)
	Print(newStack)
	Pop(newStack)
	Print(newStack)

	// Using Queue struct
	fmt.Println("Queue")
	newQueue := queue.NewQueue()
	newQueue.Push(1)
	newQueue.Push(2)
	newQueue.Push(3)
	newQueue.Push(4)
	Print(newQueue)
	Pop(newQueue)
	Print(newQueue)

	// using DS interface to call stack
	ds := NewDS("stack")
	ds.Push(1)
	ds.Push(2)
	ds.Push(3)
	ds.Push(4)
	ds.Print()
	ds.Pop()
	ds.Print()
}
