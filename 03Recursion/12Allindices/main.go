package main

import "fmt"

func main() {

	arr := []int{3, 8, 1, 8, 8, 4}
	si := 0
	data := 8
	count := 0
	ans := ALLindices(arr, si, data, count)
	fmt.Println("ans:::::", ans)

}

func ALLindices(arr []int, si, data, count int) []int {
	var indices []int

	if si == len(arr) {
		base := make([]int, count)
		return base
	}

	if arr[si] == data {
		indices = append(indices, ALLindices(arr, si+1, data, count+1)...)
	} else {
		indices = append(indices, ALLindices(arr, si+1, data, count)...)
	}

	if arr[si] == data {
		indices[count] = si
	}
	return indices
}
