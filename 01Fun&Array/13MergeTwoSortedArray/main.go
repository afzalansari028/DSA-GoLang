package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{6, 7, 8, 9, 10, 11}

	m := 3
	n := len(arr2)

	MergeTwoSortedArrays(arr1, m, arr2, n)

}

func MergeTwoSortedArrays(arr1 []int, m int, arr2 []int, n int) {

	i, j := 0, 0

	var arr3 []int
	for i < m && j < n {

		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	//first array elements remained
	for i < m {
		arr3 = append(arr3, arr1[i])
		i++
	}

	//second array elements remained
	for j < n {
		arr3 = append(arr3, arr2[j])
		j++
	}
	fmt.Println("arr3:::", arr3)
}
