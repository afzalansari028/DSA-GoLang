package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	stack := &mystack.Stack{}

	stack.Push(5)
	stack.Push(-2)
	stack.Push(9)
	stack.Push(-7)
	stack.Push(3)

	stack.Display()
	fmt.Println("-----------")
	SortStack(stack)
	stack.Display()

}

func SortStack(stack *mystack.Stack) {
	// base case
	if stack.IsEmpty() {
		return
	}

	num := stack.Peek()
	stack.Pop()

	//recursive call
	SortStack(stack)

	InsertSorted(stack, num)

}

func InsertSorted(stack *mystack.Stack, num interface{}) {

	if stack.IsEmpty() || stack.Peek().(int) < num.(int) {
		stack.Push(num)
		return
	}

	n := stack.Peek()
	stack.Pop()

	//recursive call
	InsertSorted(stack, num)
	fmt.Println("n--", n)
	stack.Push(n)

}
