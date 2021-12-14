package _365__How_Many_Numbers_Are_Smaller_Than_the_Current_Number

import (
	"container/list"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	counterList := make(map[int]*list.List, len(nums))
	for i, n := range nums {
		l := counterList[n]
		if l == nil {
			l = new(list.List)
			l.PushFront(i)
			counterList[n] = l
		} else {
			l.PushFront(i)
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	counter := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		c := len(nums) - i - counterList[nums[i]].Len()
		e := counterList[nums[i]].Front()
		counter[e.Value.(int)] = c
		counterList[nums[i]].Remove(e)
	}

	return counter
}
