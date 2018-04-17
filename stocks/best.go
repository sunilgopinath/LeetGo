package stocks

// MaxProfit returns the max profit after one buy/sell
func MaxProfit(prices []int) int {
	if len(prices) <= 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		potentialProfit := prices[i] - minPrice
		maxProfit = max(maxProfit, potentialProfit)
		minPrice = min(prices[i], minPrice)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
