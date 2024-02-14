package main

import "fmt"

func main() {

	ans := Power(2, 3)
	fmt.Println(ans)

}

func Power(a, n int) int {

	if n == 0 {
		return 1
	}
	pnm1 := Power(a, n-1)
	pn := a * pnm1
	return pn
}
