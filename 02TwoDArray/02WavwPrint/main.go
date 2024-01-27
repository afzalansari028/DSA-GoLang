package main

import "fmt"

func main() {
	arr := [][]int{ //you can give size 3x3
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(arr); i++ {

		if i%2 == 0 {
			for j := 0; j < len(arr[i]); j++ {
				fmt.Print(arr[i][j], " ")
			}
		} else {
			for j := len(arr[i]) - 1; j >= 0; j-- {
				fmt.Print(arr[i][j], " ")
			}
		}
		// fmt.Println("")
	}
}
