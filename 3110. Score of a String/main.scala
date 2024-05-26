object Solution {
    def scoreOfString(s: String): Int = {
        var sum: Int = 0
        for (i <- 1 to s.length-1) {
            sum += (s.charAt(i-1) - s.charAt(i)).abs
        }
        return sum
    }
}