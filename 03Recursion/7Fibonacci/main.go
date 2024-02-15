package main

import "fmt"

func main() {
	ans := Fibonacii(3)
	fmt.Println("ans::", ans)
}

func Fibonacii(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	Fnm1 := Fibonacii(n - 1)
	Fnm2 := Fibonacii(n - 2)

	Fib := Fnm1 + Fnm2
	return Fib
}
