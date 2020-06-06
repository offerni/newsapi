package newsapi

import (
	"testing"
)

// Assertion functions for unit tests
func assertEqual(t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("Incorrect Equals Assertion, expected: (%s), got: (%s)", expected, result)
	}
}