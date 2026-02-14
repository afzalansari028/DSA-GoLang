package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 3, 4}
	arr2 := []int{3, 4, 4, 5, 6}

	Intersection(arr1, arr2)
}

func Intersection(arr1 []int, arr2 []int) {

	set := make(map[int]bool)
	result := []int{}
	for _, val := range arr1 {
		set[val] = true
	}

	for _, val := range arr2 {

		if set[val] == true {
			result = append(result, val)
			delete(set, val) // delete to protect from duplicate
		}
	}
	fmt.Println("result::", result)
}
