package main

import "fmt"

func main() {
	ans := Factorial(4)
	fmt.Println(ans)
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	fact := n * Factorial(n-1)
	return fact
}
