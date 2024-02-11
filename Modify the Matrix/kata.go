package kata


func modifiedMatrix(matrix [][]int) [][]int {
    var maxPerCol = make([]int, len(matrix[0]))
    for i := range maxPerCol {
        maxPerCol[i] = int(math.MinInt64)
    }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == -1 {
                if maxPerCol[j] == int(math.MinInt64) {
                   currentMax := 0
                   for k := 0; k < len(matrix); k++ {
                       if matrix[k][j] > currentMax {
                           currentMax = matrix[k][j]
                       }
                   } 

                   matrix[i][j] = currentMax
                   maxPerCol[j] = currentMax

                } else {
                    matrix[i][j] = maxPerCol[j]
                }
            }
        }
    }
    return matrix
}