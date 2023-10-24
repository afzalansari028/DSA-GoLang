package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{2, 1, 4, 8, 3, 9}

	res := MaxInArray(arr)
	fmt.Println("max :", res)

	res = MinInArray(arr)
	fmt.Println("min :", res)

}

func MaxInArray(arr []int) int {

	max := math.MinInt //assignin very least value

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func MinInArray(arr []int) int {

	min := math.MaxInt //assignin very max value or you can take a value from array itself initially

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
