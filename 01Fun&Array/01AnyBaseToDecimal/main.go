package main

import "fmt"

func main() {

	srcNum := 1423
	srcBase := 5

	ans := AnyBaseToDecimal(srcNum, srcBase)
	fmt.Println("Decimal is :", ans)

}

// Divide by dest & Multiply by src
func AnyBaseToDecimal(snum int, sbase int) int {
	var ans int
	var multiplier = 1
	for snum != 0 {
		rem := snum % 10
		ans = ans + (rem * multiplier)
		multiplier = multiplier * sbase
		snum = snum / 10
	}
	return ans
}
