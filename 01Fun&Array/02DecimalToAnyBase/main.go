package main

import "fmt"

func main() {

	srcNum := 238
	destBase := 8

	ans := DecimalToAnyBase(srcNum, destBase)
	fmt.Println("dest base 5 is :", ans)

}

// Divide by dest & Multiply by src
func DecimalToAnyBase(snum int, destBase int) int {
	var ans int
	var multiplier = 1
	for snum != 0 {
		rem := snum % destBase
		ans = ans + (rem * multiplier)
		multiplier = multiplier * 10
		snum = snum / destBase
	}
	return ans
}
