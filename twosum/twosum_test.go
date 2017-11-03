package twosum_test

import (
	"testing"

	"github.com/sunilgopinath/twosum"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4}, 5, []int{1, 2}},
		{[]int{1, 2, 2, 3}, 5, []int{2, 3}},
		{[]int{1, 2, 3, 4}, 0, nil},
	}
	for _, c := range cases {
		got := twosum.TwoSum(c.in, c.target)
		if !testEq(got, c.want) {
			t.Errorf("TwoSum(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func testEq(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
