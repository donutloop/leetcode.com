package kata

func minOperations(nums []int, k int) int {
    collection := make(map[int]bool)
    var counter int
    for i := len(nums)-1; i >= 0 ; i-- {
        if len(collection) == k {
            break
        }

        if nums[i] <= k {
            v := collection[nums[i]]
            if v == false {
                collection[nums[i]] = true
            }
        }
        counter++
    }
    return counter
}