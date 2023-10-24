package main

import "fmt"

func main() {

	//taking input from user
	// var size int
	// fmt.Scan(&size)
	// arr := make([]int, size)
	// for i := 0; i < size; i++ {
	// 	fmt.Scan(&arr[i])
	// }
	// fmt.Println("arrrr", arr)

	arr := []int{10, 20, 30, 40, 50, 60}
	find := 20

	ans := LinearSearch(arr, find)
	fmt.Println("Found at index", ans)

}
func LinearSearch(arr []int, find int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == find {
			return i
		}
	}
	return -1
}
