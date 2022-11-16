package main

import "fmt"

// 5 -1 -9 8 0 1 -4 5

func maxSumAllwasIncludingRightElement(a []int) int {
	maxSum := a[len(a)-1]
	sumOfAll := a[len(a)-1]
	for i := len(a) - 2; i > 0; i-- {
		sumOfAll += a[i]
		if sumOfAll > maxSum {
			maxSum = sumOfAll
		}
	}
	return maxSum
}

func maxSumAllwasIncludingLeftElement(a []int) int {
	maxSum := a[0]
	sumOfAll := a[0]
	for i := 1; i < len(a); i++ {
		sumOfAll += a[i]
		if sumOfAll > maxSum {
			maxSum = sumOfAll
		}
	}
	return maxSum
}

func maxSumOfSubArray(a []int) int {
	if len(a) == 0 {
		panic("must not happen")
	}
	if len(a) == 1 {
		return a[0]
	}
	left, right := a[:len(a)/2], a[len(a)/2:]
	leftMaxSumIncludingRightElement := maxSumAllwasIncludingRightElement(left)
	rightMaxSumIncludingLeftElement := maxSumAllwasIncludingLeftElement(right)
	leftMaxSum := maxSumOfSubArray(left)
	rightMaxSum := maxSumOfSubArray(right)
	//fmt.Println(left, right, leftMaxSumIncludingRightElement, rightMaxSumIncludingLeftElement, leftMaxSum, rightMaxSum)
	if leftMaxSumIncludingRightElement+rightMaxSumIncludingLeftElement > leftMaxSum {
		if leftMaxSumIncludingRightElement+rightMaxSumIncludingLeftElement > rightMaxSum {
			return leftMaxSumIncludingRightElement + rightMaxSumIncludingLeftElement
		}
		return rightMaxSum
	}
	if leftMaxSum > rightMaxSum {
		return leftMaxSum
	}
	return rightMaxSum
}

func main() {
	fmt.Println(maxSumOfSubArray([]int{1, -1, -1, 2}))
}
