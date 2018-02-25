package roman

import (
	"fmt"
	"strings"
)

// IntToRoman takes a decimal input and turns it into a Roman Numeral (limit 3999)
func IntToRoman(num int) string {
	var res []string
	x := 1000
	for num > 0 {
		t := num / x
		num %= x
		res = handleDigits(t, x, res)
		x /= 10
	}

	return strings.Join(res, "")
}

func handleDigits(t, x int, curr []string) []string {
	if x == 1000 {
		// highest is 3
		curr = append(curr, strings.Repeat("M", t))
	} else if x == 100 {
		// many cases
		curr = handleUnits(t, "C", "D", "M", curr)
	} else if x == 10 {
		curr = handleUnits(t, "X", "L", "C", curr)
	} else {
		// many cases again
		curr = handleUnits(t, "I", "V", "X", curr)
	}
	return curr
}

func handleUnits(t int, c, h, a string, curr []string) []string {
	if t < 4 {
		curr = append(curr, strings.Repeat(c, t))
	} else if t == 4 {
		curr = append(curr, c+h)
	} else if t > 4 && t < 9 {
		curr = append(curr, h+strings.Repeat(c, t%5))
	} else {
		curr = append(curr, c+a)
	}
	return curr
}
