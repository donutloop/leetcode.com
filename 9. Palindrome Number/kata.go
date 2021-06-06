package ___Palindrome_Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0)
	for x > 0 {
		k := x % 10
		x = x / 10
		nums = append(nums, k)
	}

	j := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		if nums[j] != nums[i] {
			return false
		}
		j--
	}

	return true
}