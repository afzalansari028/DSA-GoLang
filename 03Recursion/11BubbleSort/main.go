package main

import "fmt"

func main() {
	arr := []int{50, 30, 20, 40, 10}
	si := 0
	li := len(arr) - 1

	//Bubble sort using recursion
	BubbleSort(arr, si, li)

}

func BubbleSort(arr []int, si, li int) {

	if li == 0 {
		return
	}

	if si == li {
		BubbleSort(arr, 0, li-1)
		return
	}

	if arr[si] > arr[si+1] {
		temp := arr[si]
		arr[si] = arr[si+1]
		arr[si+1] = temp
	}

	BubbleSort(arr, si+1, li)

	fmt.Println("arr::", arr)

}
