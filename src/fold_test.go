package main

import "testing"

func TestFold(t *testing.T) {
	cases := map[string]uint16{
		"ok":          218,
		"hello world": 1116,
	}
	for in, want := range cases {
		if got := Fold([]byte(in)); got != want {
			t.Errorf("Fold(%q) = %d, want %d", in, got, want)
		}
	}
}
