package product_test

import (
	"reflect"
	"testing"

	"github.com/sunilgopinath/LeetGo/product"
)

var testCases = []struct {
	in       []int
	expected []int
}{
	{
		[]int{1, 7, 3, 4},
		[]int{84, 12, 28, 21},
	},
}

func TestProductExceptSelf(t *testing.T) {
	for _, test := range testCases {
		got := product.ExceptSelf(test.in)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Got %v, expected %v from %v", got, test.expected, test.in)
		}
	}
}
