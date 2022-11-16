package main

import (
	"fmt"
	"math"
)

//

// Returns subarray with the largest element sum
func mss(s []int) int {
	if len(s) == 1 {
		return s[0]
	}

	l := mss(s[:len(s)/2])
	r := mss(s[len(s)/2:])

	//NOTE: MSS - maximal subarray sum
	var ml int = math.MinInt // ml stands for "middle left" - MSS inside r containing l's rightmost el
	var mr int = math.MinInt // mr statnds for "middle right" -- MSS inside l containing r's leftmost el

	for curSum, i := 0, (len(s)/2)-1; i >= 0; i-- { // calculating ml
		curSum += s[i]
		if curSum > ml {
			ml = curSum
		}
	}
	for curSum, i := 0, len(s)/2; i < len(s); i++ { // calculating mr
		curSum += s[i]
		if curSum > mr {
			mr = curSum
		}
	}

	m := ml + mr // m - mss in bounds of l and r

	// Returning maximal out of l, r and m
	if l > r {
		if l > m {
			return l
		} else {
			return m
		}
	} else /*r > l*/ {
		if r > m {
			return r
		} else {
			return m
		}
	}
}

func main() {
	//s := []int{-1, 1, 2}
	//s := []int{4, -1, 2, 1}
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(mss(s))
}
