package main

import "fmt"

func main() {
	arr := []int{10, 60, 67, 30, 20, 25}

	ans := SelectionSorting(arr)
	fmt.Println(ans)
}

func SelectionSorting(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[min] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	return arr
}
