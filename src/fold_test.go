package main

import (
	"bytes"
	"testing"
)

func TestFold(t *testing.T) {
	cases := map[string]uint16{
		"":            0,
		"A":           65,
		"ok":          218,
		"hello world": 1116,
	}
	for in, want := range cases {
		if got := Fold([]byte(in)); got != want {
			t.Errorf("Fold(%q) = %d, want %d", in, got, want)
		}
	}

	// 258 * 0xff = 65790, which exceeds 65535 and must wrap to 254.
	wrap := bytes.Repeat([]byte{0xff}, 258)
	if got := Fold(wrap); got != 254 {
		t.Errorf("Fold(258 x 0xff) = %d, want 254", got)
	}
}
