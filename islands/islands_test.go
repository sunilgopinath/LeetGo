package islands_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/islands"
)

var testCases = []struct {
	grid     [][]byte
	expected int
}{
	{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, 1},
	{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, 3},
}

func TestCount(t *testing.T) {
	for _, test := range testCases {
		got := islands.Count(test.grid)
		if got != test.expected {
			t.Fatalf("Got islands size of %d, want %d", got, test.expected)
		}
	}

}
