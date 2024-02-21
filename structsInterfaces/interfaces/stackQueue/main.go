package main

import (
	"fmt"

	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/stackQueue/queue"
	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/stackQueue/stack"
)

type Popper interface {
	Pop() int
	Print()
}

func Pop(p Popper) {
	fmt.Println("POP: ", p.Pop())
}
func Print(p Popper) {
	p.Print()
}

//implement an integer stack which satisfies the popper interface.
//if stack is empty return -1.
//you can only use PopThree function to pop the elements.
//you cannot modify the interface and function.

func main() {
	newStack := &stack.Stack{}
	newStack.Push(1)
	newStack.Push(2)
	// newStack.Push(3)
	// newStack.Push(4)
	Print(newStack)
	Pop(newStack)
	Print(newStack)

	newQueue := &queue.Queue{}
	newQueue.Push(1)
	newQueue.Push(2)
	// newQueue.Push(3)
	// newQueue.Push(4)
	Print(newQueue)
	Pop(newQueue)
	Print(newQueue)
}

//implement an integer stack which satisfies the popper interface.
//if stack is empty return -1.
//you can only use PopThree function to pop the elements.
//you cannot modify the interface and function.
