package main

import "fmt"

func main() {

	arr := []int{10, 30, 50, 20, 40}
	BubbleSort(arr)

}

func BubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-1; j++ {

			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
