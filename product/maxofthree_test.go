package product_test

import (
	"testing"

	"github.com/sunilgopinath/LeetGo/product"
)

var testThreeCases = []struct {
	in       []int
	expected int
}{
	{
		[]int{1, 2, 3},
		6,
	},
	{
		[]int{1, 2, 3, 4},
		24,
	},
	{
		[]int{-1000, -1000, -1000},
		-1000000000,
	},
	{
		[]int{-4, -3, -2, -1, 60},
		720,
	},
}

func TestMaximumProduct(t *testing.T) {
	for _, test := range testThreeCases {
		got := product.MaximumProduct(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, expected %d from %v", got, test.expected, test.in)
		}
	}
}
