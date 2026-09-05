func maxProfit(prices []int) int {
	logs := []int{}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j <= len(prices) - 1; j++ {
			profit := prices[j] - prices[i]
			logs = append(logs, profit)
		}
	}

	var maxProfit int
	for _, profit := range logs {
		maxProfit = max(profit, maxProfit)
	}

	return maxProfit
}
