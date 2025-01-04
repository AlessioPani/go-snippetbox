package assert

import "testing"

// Equal is a test helper used to assert if two comparable variables are equals.
func Equal[T comparable](t *testing.T, name string, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("Test %s | got %v, want %v", name, expected, actual)
	}
}
