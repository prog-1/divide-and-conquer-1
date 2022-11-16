package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxSubarraySum(s []int) int {
	if len(s) == 0 {
		panic("cannot happen")
	}
	if len(s) == 1 {
		return s[0]
	}
	l, r := s[:len(s)/2], s[len(s)/2:]
	var maxLeft, maxRight int
	for i, sum := len(l)-1, 0; i >= 0; i-- {
		sum += l[i]
		if maxLeft < sum {
			maxLeft = sum
		}
	}
	for i, sum := 0, 0; i < len(r); i++ {
		sum += r[i]
		if maxRight < sum {
			maxRight = sum
		}
	}
	return max(max(findMaxSubarraySum(l), findMaxSubarraySum(r)), maxLeft+maxRight)
}
