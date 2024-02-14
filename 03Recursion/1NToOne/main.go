package main

import "fmt"

func main() {
	NToOne(5) //print 5 4 3 2 1
}

func NToOne(n int) {

	if n == 0 {
		return
	}
	fmt.Println(n)
	NToOne(n - 1)
}
