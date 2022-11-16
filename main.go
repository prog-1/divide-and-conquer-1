package main

import "fmt"

func sumFinder(n []int) (maxsum int) {
	if len(n) == 0 {
		panic("no")
	}
	if len(n) == 1 {
		return n[0]
	}
	a, b := devide(n)
	var maxleft, maxright int
	for i, sum := len(a)-1, 0; i >= 0; i-- {
		sum += a[i]
		if sum > maxleft {
			maxleft = sum
		}
	}
	for i, sum := 0, 0; i < len(b); i++ {
		sum += b[i]
		if sum > maxright {
			maxright = sum
		}
	}
	return max(max(sumFinder(a), sumFinder(b)), maxleft+maxright)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func devide(n []int) (a, b []int) {
	return n[:len(n)/2], n[len(n)/2:]
}

func main() {
	ints := []int{2, 8, -2, -9, 1}
	fmt.Println(ints)
	fmt.Print(sumFinder(ints))
}
