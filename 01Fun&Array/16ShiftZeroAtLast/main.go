package main

import "fmt"

func main() {
	arr := [6]int{0, 1, 4, 3, 0, 2}

	ZeroAtlast(arr[:])

}

// bruite force
func ZeroAtlast(arr []int) {

	size := len(arr)
	result := make([]int, size)
	k := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			result[k] = arr[i]
			k++
		}
	}

	for i := k; i < len(arr); i++ {
		result[k] = 0
		k++
	}
	fmt.Println("result::", result)
}
