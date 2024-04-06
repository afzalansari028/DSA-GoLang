package main

import (
	"fmt"
)

func main() {

	result := GetPermutation("abc")
	fmt.Println(result)

}

func GetPermutation(str string) []string {

	if len(str) == 0 {
		var baseResult []string
		baseResult = append(baseResult, " ")
		return baseResult
	}

	cc := str[:1]
	ros := str[1:]

	rr := GetPermutation(ros)
	var mr []string

	for _, val := range rr {

		for j := 0; j < len(val); j++ {
			value := val[:j] + cc + val[j:]
			mr = append(mr, value)
		}
	}
	return mr
}
