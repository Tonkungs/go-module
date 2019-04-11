package main

import (
	"testing"

	"test-main/testst"
)

func TestHdplayer(t *testing.T) {
	want := "test"
	if got := testst.Testst(want); got != want {
		t.Errorf("Testst() = %q, want %q", got, want)
	}
}
