package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	stack := mystack.Stack{}

	for i := 1; i <= 5; i++ {
		stack.Push(10 * i)
	}

	stack.Display()
	fmt.Println("------------")
	ReverseStack(&stack)
	stack.Display()

}

func ReverseStack(stack *mystack.Stack) {
	//base case
	if stack.IsEmpty() {
		return
	}

	num := stack.Peek()
	stack.Pop()

	//recuscive call
	ReverseStack(stack)

	InsertAtBottom(stack, num)
}

func InsertAtBottom(stack *mystack.Stack, num interface{}) {

	// base case
	if stack.IsEmpty() {
		stack.Push(num)
		return
	}

	PoppedEle := stack.Peek()
	stack.Pop()

	// recursive call
	InsertAtBottom(stack, num)

	stack.Push(PoppedEle)

}
