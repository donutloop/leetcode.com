package kata

func finalPrices(prices []int) []int {

	var i int
	for j := 1; j < len(prices); j++ {
		if prices[j] <= prices[i] {
			prices[i] = prices[i] - prices[j]
			i++
			j = i
		} else if j == len(prices)-1 && j-i != 1 {
			i++
			j = i
		}
	}
	return prices
}
