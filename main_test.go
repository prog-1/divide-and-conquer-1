package main

import "testing"

func TestMaxSubSum(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"zero", []int{0}, 0},
		{"multiply zeros", []int{0, 0, 0, 0, 0}, 0},
		{"one number", []int{8348968438452}, 8348968438452},
		{"positive", []int{1, 1, 2, 3}, 7},
		{"negative", []int{-1, -4, -3, -0}, 4}, // ???
		{"duplicate", []int{11, 11, 11, 11, 11}, 55},
		{"both", []int{-1, 1, 2, -1}, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxSubSum(tc.s); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
