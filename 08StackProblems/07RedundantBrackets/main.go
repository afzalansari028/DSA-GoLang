package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	str := "(a+(a+b))"

	fmt.Println("IsReduntantBrackets:::", IsReduntantBrackets(str))

}

func IsReduntantBrackets(str string) bool {

	stack := mystack.Stack{}

	for i := 0; i < len(str); i++ {

		ch := string(str[i])

		if ch == "(" || ch == "+" || ch == "-" || ch == "*" || ch == "/" {
			stack.Push(ch)
		} else {
			// ")" or lower case letters
			if ch == ")" {
				IsRedundant := true

				for stack.Peek() != "(" {
					top := stack.Peek()
					if top == "+" || top == "-" || top == "*" || top == "/" {
						IsRedundant = false
					}
					stack.Pop()
				}
				if IsRedundant {
					return true
				}
				stack.Pop()
			}
		}
	}
	return false
}
