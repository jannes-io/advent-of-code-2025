package testlib

import (
	"testing"
)

func Assert[V comparable](t *testing.T, value, expected V) {
	t.Helper()

	if value != expected {
		t.Errorf(`Assertion Failed:
Expected: %v
Received: %v`, expected, value)
	}
}
