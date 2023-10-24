package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}
	find := 60
	ans := BinarySearch(arr, find)

	fmt.Println("Found at index :", ans)
}

func BinarySearch(arr []int, find int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)

		if arr[mid] == find {
			return mid
		} else if find > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
