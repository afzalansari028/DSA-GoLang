package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sl = sl[2:6]

	sl[0] = 20

	fmt.Println("val::", sl)
	fmt.Println("sl::", sl)
}
