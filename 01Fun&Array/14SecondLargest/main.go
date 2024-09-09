package main

import "fmt"

func main() {
	// second max number without sorting...
	arr := []int{5, 20, 120, 7, 10, 30}
	SecondLargest(arr)
}

func SecondLargest(arr []int) {

	max := arr[0]
	secMax := arr[0]

	//loop for max
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("max:", max)
	// loop for second max
	for i := 1; i < len(arr); i++ {
		if arr[i] > secMax && arr[i] < max {
			secMax = arr[i]
		}
	}
	fmt.Println("second max:", secMax)
}
