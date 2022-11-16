package main

import "testing"

func TestMSS(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1, -1}, 1},
		{[]int{-1, 1, 2, -1}, 3},
		{[]int{-1, 1, 2}, 3},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	} {
		got := find(tc.s)
		if got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
