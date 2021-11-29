package util

import "testing"

func AssumeSame(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q is not equal to want %q", got, want)
	}
}
