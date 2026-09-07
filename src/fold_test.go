package main

import "testing"

func TestFold(t *testing.T) {
	got := Fold([]byte("ok"))
	if got != 218 {
		t.Fatalf("Fold(ok) = %d, want 218", got)
	}
}
