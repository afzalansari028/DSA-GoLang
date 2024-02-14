package main

import "fmt"

func main() {
	PDI(5)

}

func PDI(n int) {

	if n == 0 {
		return
	}
	if n%2 != 0 {
		fmt.Println(n)
	}
	PDI(n - 1)
	if n%2 == 0 {
		fmt.Println(n)
	}
}

// func searchInsert(nums []int, target int) int {
// 	var ans int
// 	var i int

// 	if len(nums) == 0 || target <= nums[0] {
// 		ans = 0
// 		return ans
// 	}

// 	for i = 0; i < len(nums); i++ {
// 		fmt.Println(i)
// 		index := i
// 		if nums[i] == target {
// 			return i
// 		}
// 		if target > nums[index] && target < nums[index+1] {
// 			return i + 1
// 		}
// 	}
// 	fmt.Println("32")
// 	if i == len(nums)-1 {
// 		ans = len(nums)
// 		return ans
// 	}

// 	return ans
// }
