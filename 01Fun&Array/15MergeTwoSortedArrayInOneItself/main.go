package main

import "fmt"

func main() {
	arr1 := [6]int{2, 3, 6} //    2 2 3 6 7
	arr2 := [3]int{1, 2, 7}

	MergetwoArray(arr1, arr2)
}

func MergetwoArray(arr1 [6]int, arr2 [3]int) {

	p := len(arr2) - 1 //2 1 0
	q := len(arr2) - 1 //2 1 0

	k := len(arr1) - 1 //5 4 3 2 1

	for p >= 0 && q >= 0 {

		if arr1[p] > arr2[q] {
			arr1[k] = arr1[p]
			p--
			k--
		} else {
			arr1[k] = arr2[q]
			q--
			k--
		}
	}

	// for remaining elements in arr2
	for q >= 0 {
		arr1[k] = arr2[q]
		q--
		k--
	}
	fmt.Println("arr1:", arr1)
}
