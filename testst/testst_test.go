package testst

import "testing"

func TestHdplayer(t *testing.T) {
	want := "test"

	if got := Testst(want); got != want {
		t.Errorf("Testst() = %q, want %q", got, want)
	}
}
