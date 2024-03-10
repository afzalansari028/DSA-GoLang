package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := []int{9}

	PlusOne(arr)

}
func PlusOne(arr []int) {

	if len(arr) == 1 && arr[len(arr)-1]+1 > 9 {
		var arr1 [2]int
		val := arr[0]
		ans := val + 1
		ansString := strconv.Itoa(ans)

		arr1[0] = int(ansString[0] - '0')
		arr1[1] = int(ansString[1] - '0')

		fmt.Println("arr1::", arr1)
	}

	if arr[len(arr)-1]+1 > 9 {
		arr[len(arr)-1] = 0
		arr[len(arr)-2] = arr[len(arr)-2] + 1
	} else {
		arr[len(arr)-1] = arr[len(arr)-1] + 1
	}

	fmt.Println("arr::", arr)

}
