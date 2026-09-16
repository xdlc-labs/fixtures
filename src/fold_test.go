package main

import "testing"

func TestFold(t *testing.T) {
	got := Fold([]byte{200, 200})
	if got != 400 {
		t.Fatalf("Fold([200 200]) = %d, want 400", got)
	}
}
