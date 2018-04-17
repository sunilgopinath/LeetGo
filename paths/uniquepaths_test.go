package paths_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/paths"
)

var testCases = []struct {
	m        int
	n        int
	expected int
}{
	{
		7,
		3,
		28,
	},
}

func TestUniquePaths(t *testing.T) {
	for _, test := range testCases {
		got := paths.Unique(test.m, test.n)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %d, %d", got, test.expected, test.m, test.n)
		}
	}
}
