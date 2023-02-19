class Solution {
    public int minMaxDifference(int n) {
        var digitCount = ((int) Math.log10((double) n));
        var minSum = 0;
        var maxSum = 0;
        Integer maxDigit = null;
        Integer minDigit = null;
        for (; digitCount >= 0; digitCount--) {
            var base = (int) Math.pow(10, (double) digitCount);

            var digit = (n/base) % 10;

            if (maxDigit == null && digit < 9) {
                maxDigit = digit;
            }

            if(maxDigit != null && maxDigit.intValue() == digit) {
                maxSum += 9 * base;
            } else {
                maxSum += digit * base;
            }

            if (minDigit == null && digit > 0) {
                minDigit = digit;
            }

            if (minDigit != null && minDigit.intValue() == digit) {
                minSum += 0;
            } else {
                minSum += digit * base;
            }
        }
        return maxSum-minSum;
    }
}