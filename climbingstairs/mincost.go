package climbingstairs

// MinCostClimbingStairs returns the min cost of climbing stairs
func MinCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	pp := cost[0]
	p := cost[1]
	curr := 0
	for i := 2; i < len(cost); i++ {
		curr = min(pp, p) + cost[i]
		pp = p
		p = curr
	}
	return curr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
