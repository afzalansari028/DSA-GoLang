package main

import (
	"fmt"
	"strconv"
)

func main() {

	result := GetBoardPath(0, 10)

	fmt.Println(result)
	fmt.Println(len(result))

}

func GetBoardPath(curr int, end int) []string {

	if curr == end {
		br := []string{""}
		return br
	}
	if curr > end {
		br := []string{}
		return br
	}

	var mr []string
	var rr []string
	for dice := 1; dice <= 6; dice++ {
		rr = GetBoardPath(curr+dice, end)

		for _, rrs := range rr {

			mr = append(mr, strconv.Itoa(dice)+rrs)
		}

	}
	return mr
}
