package kata

func maxFrequencyElements(nums []int) int {
    var max = 0
    var total = 0
    frequencies := make(map[int]int, 0)
    for _, num := range nums {
        frequency := frequencies[num]
        frequency++
        if frequency > max {
            max = frequency
            total = 0
        }
        if frequency == max {
           total += max
        }
        frequencies[num] = frequency
    }

    return total
}