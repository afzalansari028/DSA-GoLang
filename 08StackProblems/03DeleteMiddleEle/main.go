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

	count := 0
	size := InputStack.Size()

	InputStack.Display()
	fmt.Println("-------------")
	DeleteMidElement(&InputStack, count, size)
	InputStack.Display()
}

func DeleteMidElement(InputStack *mystack.Stack, count, size int) {
	if count == size/2 {
		InputStack.Pop()
		return
	}

	num := InputStack.Peek()
	InputStack.Pop()

	//recursive call
	DeleteMidElement(InputStack, count+1, size)

	InputStack.Push(num)

}
