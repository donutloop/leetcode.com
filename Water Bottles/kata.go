package kata

func numWaterBottles(numBottles int, numExchange int) int {
	max := numBottles

	for numExchange <= numBottles {
		numBottles = numBottles - numExchange
		numBottles++
		max++
	}

	return max
}
