package main

import "fmt"

func main() {
	arr := []int{88, 11, 44, 99, 33}

	InsertionSorting(arr)
}

func InsertionSorting(arr []int) {

	for counter := 1; counter < len(arr); counter++ {

		val := arr[counter]

		j := counter - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
	fmt.Println("arr::", arr)
}
