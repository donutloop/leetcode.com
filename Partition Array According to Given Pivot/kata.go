package kata

import "container/list"

func pivotArray(nums []int, pivot int) []int {

	smaller := list.New()
	middle := list.New()
	greater := list.New()

	for _, n := range nums {
		if n < pivot {
			smaller.PushBack(n)
		} else if n == pivot {
			middle.PushBack(n)
		} else {
			greater.PushBack(n)
		}
	}

	for i := 0; i < len(nums); i++ {
		if smaller.Front() != nil {
			nums[i] = smaller.Front().Value.(int)
			smaller.Remove(smaller.Front())
		} else if middle.Front() != nil {
			nums[i] = middle.Front().Value.(int)
			middle.Remove(middle.Front())
		} else {
			nums[i] = greater.Front().Value.(int)
			greater.Remove(greater.Front())
		}
	}

	return nums
}
