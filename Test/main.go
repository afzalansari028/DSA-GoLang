package main

import "fmt"

func main() {

	arr := []int{1, 1, 2}
	fmt.Println("before:", arr)
	ass := RemoveDuplicate(arr)
	fmt.Println("asss", ass)

}

func RemoveDuplicate(arr []int) int {

	k := 0
	for i := 1; i < len(arr); i++ {
		if arr[k] != arr[i] {
			k++
			arr[k] = arr[i]
			fmt.Println("k", k)
		}
	}
	return k + 1
}
