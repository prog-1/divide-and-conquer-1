package main

import "testing"

func TestFindMaxSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want int
	}{
		{"case-1", []int{-1, 1}, 1},
		{"case-2", []int{-1, 1, 2}, 3},
		{"case-3", []int{-1, 1, 2, -1}, 3},
		{"case-4", []int{1, -1, -1, 2}, 2},
		{"case-5", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := findMaxSubarraySum(tc.s); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
