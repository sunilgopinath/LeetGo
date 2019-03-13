package number_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/number"
)

var testCases = []struct {
	in       int
	expected int
}{
	{0, 0},
	{1324, 1234},
	{4321, 1234},
	{3242, 2234},
	{1111, 1111},
	{-1111, -1111},
	{-4321, -1234},
	{3319, 1339},
	{9999999999, 9999999999},
	{-9999999999, -9999999999},
}

func TestSortInteger(t *testing.T) {
	for _, test := range testCases {
		got := number.Sort(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %d", got, test.expected, test.in)
		}
	}
}
