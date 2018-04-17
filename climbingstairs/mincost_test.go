package climbingstairs_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/climbingstairs"
)

var testCases = []struct {
	cost     []int
	expected int
}{
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for _, test := range testCases {
		got := climbingstairs.MinCostClimbingStairs(test.cost)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %v", got, test.expected, test.cost)
		}
	}
}
