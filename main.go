package main

func Maxsum(s []int, mid int) int {
	var sum, leftsum, rightsum int
	for i := mid; i >= 0; i-- {
		sum += s[i]
		leftsum = max2(leftsum, sum)
	}
	sum = 0
	for i := mid; i < len(s); i++ {
		sum += s[i]
		rightsum = max2(rightsum, sum)
	}
	return max3(leftsum+rightsum-s[mid], leftsum, rightsum)
}

func MaxSubSum(s []int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}
	if len(s) == 2 {
		return max2(s[0], s[1])
	}
	mid := (len(s) - 1) / 2
	return max3(Maxsum(s, mid-1), Maxsum(s, mid+1), Maxsum(s, mid))
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max2(max2(a, b), c)
}
