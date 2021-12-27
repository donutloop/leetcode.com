package kata

type ProductOfNumbers struct {
	numbers []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{numbers: make([]int, 0)}
}

func (this *ProductOfNumbers) Add(num int) {
	this.numbers = append(this.numbers, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.numbers) < k {
		panic("k is greater than len(this.numbers)")
	}

	sum := 1
	for _, n := range this.numbers[len(this.numbers)-k:] {
		sum = sum * n
	}
	return sum
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
