package kata

func countSeniors(details []string) int {
    var count int
    for _, detail := range details {
        var age int
        var multipler = 1
        i := len(detail)-3
        for {
            if i == 10 {
                break
            }    
            age += (int(detail[i])-48) * multipler
            multipler *= 10
            i-- 
        }
        if age > 60 {
            count++
        }
    }
    return count
}