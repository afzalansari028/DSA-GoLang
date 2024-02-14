package main

import "fmt"

func main() {
	NToN(5) //print 5 4 3 2 1 1 2 3 4 5
}

func NToN(n int) {

	if n == 0 {
		return
	}
	fmt.Println(n)
	NToN(n - 1)
	fmt.Println(n)
}
