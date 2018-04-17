package stocks_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/stocks"
)

var testCases = []struct {
	in       []int
	expected int
}{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
	{[]int{2, 4, 1}, 2},
	{[]int{2, 4, 1}, 2},
}

func TestMaxProfit(t *testing.T) {
	for _, test := range testCases {
		got := stocks.MaxProfit(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %v", got, test.expected, test.in)
		}
	}
}
