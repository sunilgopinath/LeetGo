package sortedarray_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/sortedarray"
)

var testCases = []struct {
	input    []int
	target   int
	expected int
}{
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		0,
		4,
	},
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		3,
		-1,
	},
}

var testPivotCases = []struct {
	input    []int
	expected int
}{
	{
		[]int{10, 11, 12, 2, 3, 4, 5},
		2,
	},
	{
		[]int{4, 5, 7, 0, 1, 2, 3},
		2,
	},
}

func TestSearch(t *testing.T) {
	for _, test := range testPivotCases {
		got := sortedarray.FindPivot(test.input)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %v", got, test.expected, test.input)
		}
	}
}
