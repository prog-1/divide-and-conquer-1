package main

import "fmt"

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func sumR(s []int) int {
	max, sum := s[0], s[0]
	for i := 1; i < len(s); i++ {

		if sum += s[i]; sum > max {
			max = sum
		}
	}
	return max
}

func sumL(s []int) int {
	max, sum := s[len(s)-1], s[len(s)-1]

	for i := len(s) - 2; i >= 0; i-- {
		if sum += s[i]; sum > max {
			max = sum
		}
	}
	return max
}
func lSum(s []int) int {
	if len(s) == 0 {
		panic("must not happen")
	}
	if len(s) == 1 {
		return s[0]
	}
	mid := len(s) / 2
	left, right := s[:mid], s[mid:]
	leftSum := sumL(left)
	rightSum := sumR(right)
	midSum := leftSum + rightSum

	return max(midSum, lSum(s[:mid]), lSum(s[mid:]))

}
func main() {
	fmt.Println(lSum([]int{1, -1, -1, 2}))
}
