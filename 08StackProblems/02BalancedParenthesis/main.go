package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	str := ")"

	if IsBalanced(str) {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not balanced")
	}
}

func IsBalanced(str string) bool {
	var stack mystack.Stack

	for i := 0; i < len(str); i++ {

		ch := string(str[i])
		// check for open brackets
		if ch == "[" || ch == "{" || ch == "(" {
			stack.Push(ch)
		} else {
			// for closing brackets
			if stack.IsEmpty() {
				return false
			} else {
				if ch == "]" {
					if stack.Peek() == "[" {
						stack.Pop()
					} else {
						return false
					}
				} else if ch == "}" {
					if stack.Peek() == "{" {
						stack.Pop()
					} else {
						return false
					}
				} else if ch == ")" {
					if stack.Peek() == "(" {
						stack.Pop()
					} else {
						return false
					}
				} else {
					// do nothing
				}
			}
		}
	}
	stack.Display()
	if stack.IsEmpty() {
		return true
	} else {
		return false
	}
}
