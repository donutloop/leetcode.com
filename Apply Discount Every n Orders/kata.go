package kata

type Cashier struct {
	CustomerCounter int
	n               int
	discount        float64
	products        map[int]float64
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	productsMap := make(map[int]float64, 0)
	for i, p := range products {
		productsMap[p] = float64(prices[i])
	}
	return Cashier{
		n:               n,
		discount:        float64(discount),
		products:        productsMap,
		CustomerCounter: 0,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {

	if len(product) != len(amount) {
		panic("product and amount size are unequal")
	}

	this.CustomerCounter++
	var totalPrice float64
	for i, p := range product {
		totalPrice += float64(this.products[p] * float64(amount[i]))
	}

	if this.CustomerCounter == this.n && this.discount > 0 && this.discount != 100 {
		totalPrice = totalPrice - totalPrice*(this.discount/100)
		this.CustomerCounter = 0
	} else if this.discount == 100 {
		return 0
	}

	return totalPrice
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */
