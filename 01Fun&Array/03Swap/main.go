package main

import "fmt"

func main() {

	arr := []int{10, 20, 30, 40, 50}
	swap(arr, 0, 2)

}

func swap(arr []int, i, j int) {
	fmt.Println("before swap :", arr)
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

	fmt.Println("after swap :", arr)

}
