package assert

import (
	"strings"
	"testing"
)

// Equal is a test helper used to assert if two comparable variables are equals.
func Equal[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

// StringContains is a test helper used to assert if a substring is contained in a string.
func StringContains(t *testing.T, actual string, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got %s, expected to contain %s", actual, expectedSubstring)
	}
}
