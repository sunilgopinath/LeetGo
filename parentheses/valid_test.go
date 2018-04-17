package parentheses_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/parentheses"
)

var testCases = []struct {
	in       string
	expected bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"]", false},
	{"([)]", false},
}

func TestIsValid(t *testing.T) {
	for _, test := range testCases {
		got := parentheses.IsValid(test.in)
		if got != test.expected {
			t.Fatalf("Got %v, expected %v from %s", got, test.expected, test.in)
		}
	}
}
