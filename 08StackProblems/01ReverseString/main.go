package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	str := "king"
	ans := StringReverseUsingStack(str)

	fmt.Println("ans::", ans)
}

func StringReverseUsingStack(str string) string {
	var stack mystack.Stack

	for i := 0; i < len(str); i++ {
		stack.Push(string(str[i]))
	}

	ans := ""
	for !stack.IsEmpty() {
		char := stack.Peek()
		ans = ans + char.(string)
		stack.Pop()
	}
	return ans
}
