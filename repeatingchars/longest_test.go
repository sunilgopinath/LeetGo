package repeatingchars_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/repeatingchars"
)

var testCases = []struct {
	in       string
	expected int
}{
	{
		"abcabcbb",
		3,
	},
	{
		"bbbbb",
		1,
	},
	{
		"pwwkew",
		3,
	},
	{
		"au",
		2,
	},
}

func TestUniquePaths(t *testing.T) {
	for _, test := range testCases {
		got := repeatingchars.LengthOfLongestSubstring(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %s", got, test.expected, test.in)
		}
	}
}
