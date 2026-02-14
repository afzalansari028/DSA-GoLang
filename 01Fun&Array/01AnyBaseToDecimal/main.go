package main

import (
	"fmt"
	"math"
)

func main() {

	srcNum := 1423
	srcBase := 5

	ans := AnyBaseToDecimal(srcNum, srcBase)
	fmt.Println("Decimal is :", ans)

	Power(2, 6)

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

// / alternate solution--
func AnyBaseToDecimal1(srcNum, srcBase int) int {
	res := 0
	var k float64 = 0
	srcB := float64(srcBase)
	for srcNum != 0 {
		ld := srcNum % 10
		basewithPower := math.Pow(srcB, k)
		res = res + ld*int(basewithPower)
		k++
		srcNum = srcNum / 10
	}
	return res
}

// Power function
func Power(x, y int) {

	var res int = 1
	for i := 0; i < y; i++ {
		res = res * x
	}
	fmt.Println("res--", res)
}
