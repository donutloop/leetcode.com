package kata

func singleNonDuplicate(nums []int) (a int) {
	for _, n := range nums {
		a = a ^ n
	}
	return
}
