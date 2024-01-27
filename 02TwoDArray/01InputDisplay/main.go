package main

import "fmt"

func main() {

	array := [][]int{ //you can give size 3x3
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("arr", array)

}

func Display(arr [][]int) {
	// for _, val := range arr {
	// 	fmt.Print(val)
	// }

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr[i]); j++ {

			fmt.Print(arr[i][j], " ")
		}
		fmt.Println("")
	}
}
