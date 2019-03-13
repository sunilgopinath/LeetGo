package lower_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/lower"
)

func TestToLowerCase(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in       string
		expected string
	}{
		{
			"Abc",
			"abc",
		},
	}
	for _, d := range testCases {
		got := lower.ToLowerCase(d.in)
		if got != d.expected {
			t.Fatalf("Expected %s, got %s", d.expected, got)
		}
	}
}
