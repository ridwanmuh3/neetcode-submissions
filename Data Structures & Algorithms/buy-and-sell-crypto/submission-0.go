func maxProfit(prices []int) int {
	memo := map[int]int{}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j <= len(prices) - 1; j++ {
			profit := prices[j] - prices[i]
			memo[profit]++
		}
	}

	var maxProfit int
	for p, _ := range memo {
		maxProfit = max(p, maxProfit)
	}

	return maxProfit
}
