package kata

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	begin := nums[0]
	last := -1
	counter := 0
	continuous := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			last = nums[i]
			counter = counter + 1
			if len(nums)-1 == i {
				continuous = append(continuous, strconv.Itoa(begin)+"->"+strconv.Itoa(last))
			}
			continue
		}

		if counter > 0 {
			continuous = append(continuous, strconv.Itoa(begin)+"->"+strconv.Itoa(last))
		} else {
			continuous = append(continuous, strconv.Itoa(begin))
		}

		begin = nums[i]
		counter = 0
		if len(nums)-1 == i {
			continuous = append(continuous, strconv.Itoa(nums[i]))
		}
	}
	return continuous
}
