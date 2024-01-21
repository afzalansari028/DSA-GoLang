package main

import "fmt"

func main() {
	arr := []int{88, 11, 44, 99, 33}

	InsertionSorting(arr)
}

func InsertionSorting(arr []int) {

	for i := 1; i < len(arr); i++ {

		temp := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	fmt.Println("arr::", arr)
}
