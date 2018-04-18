package product

import (
	"math"
)

// MaximumProduct returns max of 3 products
func MaximumProduct(nums []int) int {
	// keep track of highest 3
	// lowest 2 (in case of negative)
	a, b, c := math.MinInt32, math.MinInt32, math.MinInt32
	y, z := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			c = b
			b = a
			a = nums[i]
		} else if nums[i] > b {
			c = b
			b = nums[i]
		} else if nums[i] > c {
			c = nums[i]
		}
		if nums[i] < z {
			y = z
			z = nums[i]
		} else if nums[i] < y {
			y = nums[i]
		}
	}

	return max(a*b*c, z*y*a)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
