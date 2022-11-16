package main

import "testing"

func TestUnimodalMaxSum(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want int
	}{
		{"1", []int{-1, 1}, 1},
		{"2", []int{-1, 1, 2}, 3},
		{"3", []int{-1, -2, -4}, -1},
		{"4", []int{1, -1, -1, 2}, 2},
		{"5", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := unimodalMaxSum(tc.s, 0, 0); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
