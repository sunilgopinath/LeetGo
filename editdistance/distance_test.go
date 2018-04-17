package editdistance_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/editdistance"
)

var testCases = []struct {
	word1    string
	word2    string
	expected int
}{
	{
		"horse",
		"ros",
		3,
	},
	{
		"",
		"ros",
		3,
	},
}

func TestMinDistance(t *testing.T) {
	for _, test := range testCases {
		got := editdistance.MinDistance(test.word1, test.word2)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %s %s", got, test.expected, test.word1, test.word2)
		}
	}
}
