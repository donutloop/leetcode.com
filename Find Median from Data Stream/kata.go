package kata

import "sort"

type MedianFinder struct {
	list []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{list: make([]int, 0)}
}

func (this *MedianFinder) AddNum(num int) {
	this.list = append(this.list, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Slice(this.list, func(i, j int) bool {
		return this.list[i] < this.list[j]
	})
	if len(this.list)%2 == 1 {
		return float64(this.list[len(this.list)/2])
	}
	return (float64(this.list[len(this.list)/2-1]) + float64(this.list[len(this.list)/2])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
