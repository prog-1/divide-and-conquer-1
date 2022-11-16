package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func find(s []int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}
	var l, r int
	a, b := s[:len(s)/2], s[len(s)/2:]
	for i, sum := len(a)-1, 0; i >= 0; i-- {
		sum += a[i]
		if l < sum {
			l = sum
		}
	}
	for i, sum := 0, 0; i < len(b); i++ {
		sum += b[i]
		if r < sum {
			r = sum
		}
	}
	return max(max(find(a), find(b)), l+r)
}

func main() {
	//s := []int{-1, 1, 2}
	//s := []int{4, -1, 2, 1}
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(find(s))
}
