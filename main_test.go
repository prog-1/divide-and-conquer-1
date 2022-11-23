package main

import "testing"

func TestMaxSubSum(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want int
	}{
		{"1", []int{3, 5, 2, 4, 7, 3}, 24},
		{"2", []int{3, -1, 3, -4}, 5},
		{"3", []int{-2, 1, -2, 4, -1, 2, 1, -5, 4}, 6},
		{"4", []int{-6, -2, -5, -9, -6, -1, -7, -5}, -1},
		{"5", []int{1, 1, 1, 1, 1, 1}, 6},
		{"6", []int{0}, 0},
		{"7", []int{}, 0},
		{"8", nil, 0},
	} {
		got := maxSubSum(tc.s)
		if got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
