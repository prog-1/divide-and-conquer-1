package main

import (
	"fmt"
)

func maxnum(a, b, c int) (res int) {
	if a >= b && a >= c {
		res = a
	} else if b >= a && b >= c {
		res = b
	} else {
		res = c
	}
	return res
}

func unimodalMaxSum(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	mid := len(s) / 2
	sum := 0
	maxls := 0
	for i := mid; i >= 0; i-- {
		sum += s[i]
		if sum > maxls {
			maxls = sum
		}
	}
	sum = 0
	maxrs := 0
	for i := mid + 1; i < len(s); i++ {
		sum += s[i]
		if sum > maxrs {
			maxrs = sum
		}
	}
	maxmid := maxrs + maxls
	return maxnum(unimodalMaxSum(s[:mid]), unimodalMaxSum(s[mid:]), maxmid)
}

func main() {
	for _, in := range [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
	} {
		fmt.Println(unimodalMaxSum(in))
	}
}
