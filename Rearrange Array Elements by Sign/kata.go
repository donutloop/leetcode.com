package kata

import "container/list"

func rearrangeArray(nums []int) []int {

	negative := list.New()
	positive := list.New()

	for _, n := range nums {
		if n < 0 {
			negative.PushBack(n)
		} else {
			positive.PushBack(n)
		}
	}

	nums[0] = positive.Front().Value.(int)
	positive.Remove(positive.Front())

	for i := 1; i < len(nums); i++ {
		if i%2 == 1 {
			nums[i] = negative.Front().Value.(int)
			negative.Remove(negative.Front())
		} else {
			nums[i] = positive.Front().Value.(int)
			positive.Remove(positive.Front())
		}
	}

	return nums
}
