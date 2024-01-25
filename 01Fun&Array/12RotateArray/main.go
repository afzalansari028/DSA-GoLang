package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7}
	d := 2
	for i := 0; i < d; i++ {
		OneRotate(array, d)
	}
}

func OneRotate(arr []int, d int) {

	temp := arr[0]

	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
	fmt.Println("array:", arr)
}
