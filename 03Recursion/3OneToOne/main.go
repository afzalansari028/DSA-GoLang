package main

import "fmt"

func main() {
	OneToN(5)
}

func OneToN(n int) {

	if n == 0 {
		return
	}
	fmt.Println(n)
	OneToN(n - 1)
	fmt.Println(n)
}
