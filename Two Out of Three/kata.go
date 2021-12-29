package Two_Out_of_Three

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {

	type numCounter struct {
		a, b, c int
	}

	counter := make(map[int]numCounter)

	for _, num := range nums1 {
		if counter[num].a == 0 {
			counter[num] = numCounter{a: 1, b: 0, c: 0}
		}
	}

	for _, num := range nums2 {
		c := counter[num]
		if c.b == 0 {
			c.b++
			counter[num] = c
		}
	}

	for _, num := range nums3 {
		c := counter[num]
		if c.c == 0 {
			c.c++
			counter[num] = c
		}
	}

	matched := make([]int, 0)
	for num, c := range counter {
		if (c.a + c.b + c.c) >= 2 {
			matched = append(matched, num)
		}
	}

	return matched
}
