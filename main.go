package main

import "fmt"

func main() {
	s := []int{}
	fmt.Println(maxSubSum(s))
}

func maxSubSum(s []int) int {
	//base cases
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}

	//splitting array in halfs
	lS, rS := split(s)

	//taking max subslice from left and right sides
	lSS := maxSubSum(lS)
	rSS := maxSubSum(rS)

	// finding center subslice sum
	rCSS := rS[0]
	for i, sum := 0, 0; i < len(rS); i++ {
		sum += rS[i]
		if sum > rCSS {
			rCSS = sum
		}
	}

	lCSS := lS[len(lS)-1]
	for i, sum := len(lS)-1, 0; i >= 0; i-- {
		sum += lS[i]
		if sum > lCSS {
			lCSS = sum
		}
	}

	cSS := lCSS + rCSS

	//finding maximum from three
	return max(lSS, rSS, cSS)

}

func split(s []int) (lS, rS []int) {
	mid := len(s) / 2
	lS = s[:mid]
	rS = s[mid:]

	return lS, rS
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}
