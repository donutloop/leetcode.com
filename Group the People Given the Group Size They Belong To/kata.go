package kata

func groupThePeople(groupSizes []int) [][]int {
	buckets := make(map[int][]int, 0)
	dividedGroups := make([][]int, 0)
	for i, groupSize := range groupSizes {
		bucket := buckets[groupSize]
		if bucket == nil {
			bucket = make([]int, 0)
		}

		bucket = append(bucket, i)
		if len(bucket) == groupSize {
			dividedGroups = append(dividedGroups, bucket)
			buckets[groupSize] = nil
		} else {
			buckets[groupSize] = bucket
		}
	}
	return dividedGroups
}
