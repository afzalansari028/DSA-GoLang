package main

import "fmt"

func main() {
	arr := []int{6, 8, 1, 1, 8, 3, 4}
	si := 0
	data := 4

	ans := FindFirstIndex(arr, si, data)
	fmt.Println("answer::", ans)
}

func FindFirstIndex(arr []int, si, data int) int {

	if len(arr) == si {
		return -1
	}

	if arr[si] == data {
		return si
	} else {
		return FindFirstIndex(arr, si+1, data)
	}

}
