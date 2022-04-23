package Most_Frequent_Number_Following_Key_In_an_Array

func mostFrequent(nums []int, key int) int {

	counter := map[int]int{}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			counter[nums[i+1]]++
		}
	}

	var max, num int
	for n, c := range counter {
		if c > max {
			max = c
			num = n
		}
	}

	return num
}
