package merge_test

import (
	"reflect"
	"testing"

	"github.com/sunilgopinath/LeetGo/merge"
)

var testCases = []struct {
	in1      []int
	l1       int
	in2      []int
	l2       int
	expected []int
}{
	{
		[]int{1, 4, 7, 9, 0, 0, 0},
		4,
		[]int{2, 3, 8},
		3,
		[]int{1, 2, 3, 4, 7, 8, 9},
	},
	{
		[]int{0},
		0,
		[]int{1},
		1,
		[]int{1},
	},
	{
		[]int{4, 5, 6, 0, 0, 0},
		3,
		[]int{1, 2, 3},
		3,
		[]int{1, 2, 3, 4, 5, 6},
	},
}

func TestSortedArray(t *testing.T) {
	for _, test := range testCases {
		got := merge.SortedArray(test.in1, test.l1, test.in2, test.l2)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Got %v, expected %v", got, test.expected)
		}
	}
}
