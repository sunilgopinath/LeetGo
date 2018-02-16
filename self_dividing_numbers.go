// Package selfdividing finds all self dividing numbers: https://leetcode.com/problems/self-dividing-numbers/description/
package selfdividing

func selfDividingNumbers(left int, right int) []int {
    var res []int
    for i := left; i <= right; i++ {
        if isSelfDividing(i) {
            res = append(res, i)
        }
    }
    return res
}

func isSelfDividing(n int) bool {
	t := n
	for n > 0 {
        if n%10 == 0 || t%(n%10) != 0 {
            return false
        }
		n /= 10
	}
	return true
}
