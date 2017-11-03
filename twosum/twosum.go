// Package twosum contains the answer to https://leetcode.com/problems/two-sum/description/
package twosum

//TwoSum finds a pair of numbers which adds to the target
func TwoSum(nums []int, target int) []int {
	// go through list looking for target - current loop variable
	m := make(map[int]int)
	for i, e := range nums {
		if _, ok := m[target-e]; ok {
			r := []int{m[target-e], i}
			return r
		}
		m[e] = i
	}
	return nil
}
