package main

import "fmt"

func main() {

	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	GetARowValues(arr)
	GetAColValues(arr)
	SumOfElements(arr)
	Transpose(arr)
}

func GetARowValues(arr [][]int) {
	fmt.Println("row 0::", arr[0])
}
func GetAColValues(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("column 0 ::", arr[i][0])
	}
}

// sum of all elements
func SumOfElements(arr [][]int) {

	sum := 0
	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr[i]); j++ {
			sum = sum + arr[i][j]
		}
	}
	fmt.Println("elements sum:::", sum)
}

// // transpose of matrix - row -> col, col -> row
// func Transpose(arr [][]int) {

// 	for i := 0; i < len(arr); i++ {

// 		for j := 0; j < len(arr[i]); j++ {

// 			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
// 		}
// 	}

// 	fmt.Println("transposed matrix::", arr[1][0])

// }
