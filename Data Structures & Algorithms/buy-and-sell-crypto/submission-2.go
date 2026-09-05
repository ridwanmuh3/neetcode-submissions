/*func maxProfit(prices []int) int {
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
}*/

func maxProfit(prices []int) int {
	minBuy, maxPrice := math.MaxInt32, 0
	for _, sell := range prices {
		if sell < minBuy {
			minBuy = sell
		}
		if sell - minBuy > maxPrice {
			maxPrice = sell - minBuy
		}
	}
	return maxPrice
}
