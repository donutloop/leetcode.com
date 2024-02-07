package kata

func findMissingAndRepeatedValues(grid [][]int) []int {
    var gridRowCount = len(grid)
    var gridColCount = len(grid[0])
    
    var counter = make([]int, gridRowCount*gridColCount)

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            counter[grid[i][j] - 1]++
        }   
    }

    var missingNumber int
    var repeatingNumber int
    for i := 0; i < len(counter); i++ {
        if counter[i] == 0 {
            missingNumber = i+1
        } else if counter[i] == 2 {
            repeatingNumber = i+1
        }
    }   

   return []int{repeatingNumber, missingNumber}
}