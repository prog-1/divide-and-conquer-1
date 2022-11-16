package main

import "testing"

func TestUnimodalMax(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []int
		want int
	}{
		{"empty", []int{}, 0},
		{"example01", []int{-1, 1}, 1},
		{"example02", []int{-1, 1, 2}, 3},
		{"example03", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := unimodalMax(tc.in)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
