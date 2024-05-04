package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	fmt.Println("let's play with stack using golang")

	stack := mystack.Stack{}

	fmt.Println("isempty::", stack.IsEmpty())

	stack.Push(10) // push an elemment into the stack int type
	stack.Push(20)

	stack.Push("a")
	stack.Push("b")

	stack.Display()
	stack.Pop() // remomve an elememnt from the stack

	stack.Display()                            //display all elements
	fmt.Println("top element::", stack.Peek()) // see the top element

	fmt.Println("isempty::", stack.IsEmpty())
	fmt.Println("size::", stack.Size())
}
