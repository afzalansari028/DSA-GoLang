package main

import "fmt"

func main() {

	arr := []int{1, 3}
	target := 3
	res := searchInsert(arr, target)
	fmt.Println("res:::", res)
}

func searchInsert(nums []int, target int) int {
	var ans int
	var i int

	if len(nums) == 0 || target <= nums[0] {
		ans = 0
		return ans
	}

	for i = 0; i < len(nums); i++ {
		fmt.Println(i)
		index := i
		if nums[i] == target {
			return i
		}
		if target > nums[index] && target < nums[index+1] {
			return i + 1
		}
	}
	fmt.Println("32")
	if i == len(nums)-1 {
		ans = len(nums)
		return ans
	}

	return ans
}
