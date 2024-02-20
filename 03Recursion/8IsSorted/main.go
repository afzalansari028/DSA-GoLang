package main

import "fmt"

func main() {

	arr := []int{3, 8, 6, 1, 9, 7}
	// arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("IsSorted::", IsSorted(arr, 0))

}

func IsSorted(arr []int, si int) bool {

	if si == len(arr)-1 {
		return true
	}
	if arr[si] > arr[si+1] {
		return false
	}

	ans := IsSorted(arr, si+1)
	return ans

}
