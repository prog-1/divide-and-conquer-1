package main

import "testing"

func TestLSum(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want int
	}{
		{"1", []int{-1, 1}, 1},
		{"2", []int{-1, 1, 2}, 3},
		{"3", []int{1, -1, -1, 2}, 2},
		{"4", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"5", []int{1, 2, 3}, 6},
		{"6", []int{-1, 1, 2, 32}, 35},
		{"7", []int{1, -1, -1, 2}, 2},
		{"8", []int{2, -4, 1, 9, -6, 7, -3}, 11},
		{"9", []int{1}, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := lSum(tc.s); got != tc.want {
				t.Errorf("case:%v - got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}
