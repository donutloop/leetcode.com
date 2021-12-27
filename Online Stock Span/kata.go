package kata

type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{prices: make([]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	this.prices = append(this.prices, price)
	if len(this.prices) == 1 {
		return 1
	}

	var c int
	for i := len(this.prices) - 1; i >= 0; i-- {
		if this.prices[i] <= price {
			c++
			continue
		}
		break
	}
	return c
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
