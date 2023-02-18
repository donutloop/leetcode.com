class Solution {
    public int alternateDigitSum(int n) {
        var digitCount = ((int) Math.log10((double) n));
        var sum = 0;
        var sign = 1;
        for (; digitCount >= 0; digitCount--) {
            var digit = (n/(int) Math.pow(10, (double) digitCount)) % 10;
            sum += sign * digit;
            sign = sign * -1;
        }
        return sum;
    }
}