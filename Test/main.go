package main

import (
	"fmt"
	"math"
)

func main() {
	srcNum := 1011
	srcBase := 2

	result := AnyBaseToDecimal(srcNum, srcBase)
	fmt.Println("resulta::", result)

	Power(2, 6)

}

func AnyBaseToDecimal(srcNum, srcBase int) int {
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
