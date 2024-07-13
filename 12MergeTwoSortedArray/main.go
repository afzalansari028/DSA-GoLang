package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := []int{10, 30, 60, 70}
	arr2 := []int{20, 50, 80, 90, 100}

	Merge(arr1, arr2)
	MurgeTwoArr(arr1, arr2)

}

// brute force approach
func Merge(arr1, arr2 []int) {
	p, q, k := 0, 0, 0

	finalArray := make([]int, len(arr1)+len(arr2))
	for p < len(arr1) {
		finalArray[k] = arr1[p]
		p++
		k++
	}
	for q < len(arr2) {
		finalArray[k] = arr2[q]
		q++
		k++
	}

	sort.Ints(finalArray)

	fmt.Println("merged arr::", finalArray)
}

// using 3 pointer as variables
func MurgeTwoArr(arr1 []int, arr2 []int) {

	p, q := 0, 0
	k := 0

	Size := len(arr1) + len(arr2)
	finalArray := make([]int, Size)

	for p < len(arr1) && q < len(arr2) {

		if arr1[p] < arr2[q] {
			finalArray[k] = arr1[p]
			k++
			p++
		} else {
			finalArray[k] = arr2[q]
			k++
			q++
		}
	}
	// add the remaining elements int the final array
	for p < len(arr1) {

		finalArray[k] = arr1[p]
		k++
		p++
	}

	for q < len(arr2) {
		finalArray[k] = arr2[q]
		k++
		q++
	}
	fmt.Println("arr::", finalArray)

}
