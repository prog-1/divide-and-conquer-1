package main

import "testing"

func TestMSS(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"0", []int{0}, 0},
		{"multiple 0", []int{0,0,0}, 0},
		{"+", []int{1, 23, 4, 56}, 84},
		{"-", []int{-1, -43, -9}, -1},
		{"dublicates", []int{12, 12, 12}, 36},
		{"1", []int{1, -1}, 1},
		{"2", []int{-1, 1, 2}, 3},
		{"4", []int{-2, 1, -2, 4, -1, 2, 1, -5, 4}, 6},
	} {
		got := mss(tc.input)
		if got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
