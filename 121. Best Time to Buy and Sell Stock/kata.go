package _21__Best_Time_to_Buy_and_Sell_Stock

func maxProfit(prices []int) int {
	min := -1
	var maxProfit int
	for i := 0; i < len(prices); i++ {
		if min == -1 || min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
	}
	return maxProfit
}
