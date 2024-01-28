package main

import "fmt"

func main() {

	arr := [][]int{ //you can give size 3x3
		{11, 12, 13, 14},
		{21, 22, 23, 24},
		{31, 32, 33, 34},
		{41, 42, 43, 44},
	}

	SpiralPrint(arr)
}

func SpiralPrint(arr [][]int) {
	top := 0
	bottom := len(arr) - 1
	left := 0
	right := len(arr[top]) - 1
	count := (bottom + 1) * (right + 1)
	dir := 1

	for left <= right && top <= bottom {

		if count > 0 {
			if dir == 1 {
				for i := left; i <= right; i++ {
					fmt.Print(arr[top][i], " ")
					count--
				}
				dir = 2
				top++
			}
		}
		if count > 0 {
			if dir == 2 {
				for i := top; i <= bottom; i++ {
					fmt.Print(arr[i][right], " ")
					count--
				}
				dir = 3
				right--
			}
		}
		if count > 0 {
			if dir == 3 {
				for i := right; i >= left; i-- {
					fmt.Print(arr[bottom][i], " ")
					count--
				}
				dir = 4
				bottom--
			}
		}
		if count > 0 {
			if dir == 4 {
				for i := bottom; i >= top; i-- {
					fmt.Print(arr[i][left], " ")
					count--
				}
				dir = 1
				left++
			}
		}
	}
}
