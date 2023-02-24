class Solution {
    public int divisorSubstrings(int num, int k) {
        var oldNum = num;
        var digitCount = ((int) Math.log10((double) num));
        var sum = 0;
        var countK = k-1;
        var count = 0;
        for (; digitCount >= 0; digitCount--) {
            var digit = (num/(int) Math.pow(10, (double) digitCount)) % 10;

            sum += digit * Math.pow(10, (double) countK);
            countK--;
            if (countK < 0) {
                if (sum > 0 && oldNum % sum == 0) count++;
                countK = k - 1;
                if (k != 1) {
                    digitCount += k-1;
                }
                sum = 0;
            }
        }
        return count;
    }
}