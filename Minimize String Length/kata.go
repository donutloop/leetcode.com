package kata

func minimizedStringLength(s string) int {
    var minimizedLength = len(s)
    sb := []byte(s)
    sort.Slice(sb, func(i int, j int) bool { 
        return sb[i] < sb[j] 
    })
    for i := 1; i < len(sb); i++ {
        if sb[i] == sb[i-1] {
            minimizedLength--
        }
    }
    return minimizedLength
}