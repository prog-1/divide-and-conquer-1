package main

import "testing"

func TestMaxSumOfSubArray(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  int
	}{
		{[]int{-1, 1, 2, -1}, 3},
		{[]int{1, -1, -1, 2}, 2},
		{[]int{1, 2, -8, 1, 3, 9, -2, 1}, 13},
		{[]int{1, 0, 1}, 2},
	} {
		if got := maxSumOfSubArray(tc.input); tc.want != got {
			t.Errorf("maxSumofArray(%v) = %v, want = %v", tc.input, got, tc.want)
		}
	}
}
