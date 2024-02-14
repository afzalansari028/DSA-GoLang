package main

import "fmt"

func main() {
	PrintDecSkip(5) //print 5 4 1 2 4
}

func PrintDecSkip(n int) {
	if n == 0 {
		return
	}
	if n%2 == 1 {
		fmt.Println(n)
	}
	PrintDecSkip(n - 1)
	if n%2 == 0 {
		fmt.Println(n)
	}

}
