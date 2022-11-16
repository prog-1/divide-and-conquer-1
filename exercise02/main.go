package main

import (
	"math"
)

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxSum(in []int) int {
	if len(in) == 0 {
		return 0
	}
	sums := make([]int, len(in))
	sums[0] = in[0]
	for i := 1; i < len(in); i++ {
		sums[i] = max(in[i]+sums[i-1], in[i])
	}
	maxVal := math.MinInt32
	for i := 0; i < len(sums); i++ {
		maxVal = max(sums[i], maxVal)
	}
	return maxVal
}

func check(arr1, arr2, midArr []int) int {
	if maxSum(arr1) > maxSum(arr2) && maxSum(arr1) > maxSum(midArr) {
		return maxSum(arr1)
	} else if maxSum(arr2) > maxSum(arr1) && maxSum(arr2) > maxSum(midArr) {
		return maxSum(arr2)
	} else {
		return maxSum(midArr)
	}
}

func unimodalMax(in []int) int {
	arr1 := in[:len(in)/2]
	arr2 := in[len(in)/2:]
	midArr := append(arr1[len(arr1)/2:], arr2[:len(arr2)/2]...)
	return check(arr1, arr2, midArr)
}

func main() {
}
