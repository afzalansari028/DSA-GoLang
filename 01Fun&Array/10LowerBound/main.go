package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 2, 3, 3, 3, 4, 4, 9, 11}
	result := LowerBound(3, arr)
	fmt.Println("result:::", result)
}

func LowerBound(data int, arr []int) int {

	low := 0
	high := len(arr) - 1
	ans := -1

	for low <= high {
		mid := (low + high) / 2
		fmt.Println("mid:", mid)

		if data == arr[mid] {
			ans = mid
			high = mid - 1
		} else if data < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
