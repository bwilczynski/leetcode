package p121

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProfit
}
