package main

import "fmt"

func main() {
	ans := GetSubsequences("abc")
	fmt.Println("ans:::", ans)
	fmt.Println("ans:::", len(ans))
}

func GetSubsequences(str string) []string {
	// base case
	if len(str) == 0 {
		var baseResult []string
		baseResult = append(baseResult, "")
		return baseResult
	}

	cc := str[:1]
	ros := str[1:]

	var myResult []string
	var recResult = GetSubsequences(ros)

	for i := 0; i < len(recResult); i++ {
		myResult = append(myResult, recResult[i])
		myResult = append(myResult, cc+recResult[i])
	}
	return myResult
}
