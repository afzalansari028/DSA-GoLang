package main

import "fmt"

func main() {

	result := GetMMazePath(0, 0, 2, 2)
	fmt.Println("result :", result)

}

func GetMMazePath(cr int, cc int, er int, ec int) []string {

	if cr == er && cc == ec {
		br := []string{""}
		return br
	}
	if cr > er || cc > ec {
		br := []string{}
		return br
	}

	var mr []string

	rrh := GetMMazePath(cr, cc+1, er, ec)
	for _, rrhs := range rrh {
		mr = append(mr, "H"+rrhs)
	}

	rrv := GetMMazePath(cr+1, cc, er, ec)
	for _, rrvs := range rrv {
		mr = append(mr, "V"+rrvs)
	}
	return mr
}
