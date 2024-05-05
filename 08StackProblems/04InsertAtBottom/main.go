package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {
	InputStack := mystack.Stack{}

	InputStack.Push(10)
	InputStack.Push(20)
	InputStack.Push(30)
	InputStack.Push(40)
	InputStack.Push(50)

	// size := InputStack.Size()
	InputStack.Display()
	fmt.Println("-----------------")
	AddAtLast(&InputStack, 90)
	InputStack.Display()
}

func AddAtLast(InputStack *mystack.Stack, x int) {

	if InputStack.IsEmpty() {
		InputStack.Push(x)
		return
	}

	PoppedEle := InputStack.Pop()
	fmt.Println("PoppedEle::", PoppedEle)

	// recursive call
	AddAtLast(InputStack, x)

	InputStack.Push(PoppedEle)

}
