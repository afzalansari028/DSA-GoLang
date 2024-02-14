package main

import "fmt"

func main() {
	ans := Factorial(5)
	fmt.Println(ans)
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	fact_nm1 := Factorial(n - 1)

	fact := n * fact_nm1

	return fact
}
