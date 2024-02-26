package main

import "fmt"

func main() {
	arr := []int{6, 8, 1, 1, 8, 3, 4}
	si := 0
	data := 8

	ans := FindLastIndex(arr, si, data)
	fmt.Println("last index::", ans)
}

func FindLastIndex(arr []int, si, data int) int {

	if len(arr) == si {
		return -1
	}

	index := FindLastIndex(arr, si+1, data)

	if index == -1 {

		if arr[si] == data {
			return si
		} else {
			return -1
		}
	}
	return index

}
