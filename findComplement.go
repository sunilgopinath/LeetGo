// Package complement returns the complement of an integer https://leetcode.com/problems/number-complement/description/
package complement

func findComplement(num int) int {
    return toDecimal(complement(toBinary(num)))
}

func toBinary(n int) string {
	s := strconv.FormatInt(int64(n), 2)
	return s
}

func toDecimal(s string) int {
	var n float64
	for i := len(s) - 1; i >= 0; i-- {
		t, _ := strconv.Atoi(string(s[i]))
		n += float64(t) * math.Pow(2, float64(len(s)-1-i))
	}
	return int(n)
}

func complement(s string) string {
	var res string
	for _, r := range []rune(s) {
		if r-'0' == 1 {
			res += "0"
		} else {
			res += "1"
		}
	}
	return res
}
