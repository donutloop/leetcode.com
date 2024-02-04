package kata

func findWordsContaining(words []string, x byte) []int {
    var indices []int
    for i, word := range words {
        for _, char := range word {
            if byte(char) == x {
                indices = append(indices, i)
                break
            }
        }
    }
    return indices
}