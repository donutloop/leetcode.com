package kata

func sumOfMultiples(n int) int {
    sum := 0
    for i := 1; i < n+1 ; i++ {
        if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
            sum += i
        }        
    }
    return sum
}