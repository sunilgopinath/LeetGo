package number

import (
	"math"
)

// Sort returns number sorted in ascending order
func Sort(n int) int {

	// handling negative numbers
	isNegative := false
	if n < 0 {
		isNegative = true
	}
	n = abs(n)

	// first get the digits of the number
	ss := digits(n)

	// now sort the digits (returns float64 as Go returns math.Pow10 as float)
	res := sortDigits(ss)

	// restore negative if necessary
	if isNegative {
		res *= -1
	}

	return int(res)
}

func digits(n int) []int {
	// O(1) space
	ss := make([]int, 10)
	for n > 0 {
		ss[n%10]++
		n /= 10
	}
	return ss
}

func sortDigits(ss []int) float64 {
	// Even though it is a double for loop it is not O(n^2)
	// first slice is O(1)
	// second slice is O(n)
	var res float64
	var tc int
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] != 0 {
			for j := 0; j < ss[i]; j++ {
				res += float64(i) * math.Pow10(tc)
				tc++
			}

		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
