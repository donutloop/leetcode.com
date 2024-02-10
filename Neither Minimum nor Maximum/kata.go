package kata

import (
    "sort"
)

func findNonMinOrMax(nums []int) int {
    if len(nums) <= 2 {
        return -1
    }
    sort.Ints(nums)
    return nums[1]
}