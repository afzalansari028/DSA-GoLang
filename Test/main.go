package main

import "fmt"

func main() {
	str := []string{"ab", "cd"}

	for val := range str {
		// fmt.Println("i", i)
		fmt.Println("val", val)
	}
}
