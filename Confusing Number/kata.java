class Solution {
    public boolean confusingNumber(int num) {
        var digitCount = ((int) Math.log10((double) num));
        var reversedCount = 0;
        var oldNum = num;
        var newNum = 0;
        for (; digitCount >= 0; digitCount--) {
            var digit = (num/(int) Math.pow(10, (double) digitCount)) % 10;
            if (digit == 2 || digit == 3 || digit == 4 || digit == 5 || digit == 7) {
                return false;
            }
            if (digit == 6) {
                digit = 9;
            } else if (digit == 9) {
                digit = 6;
            }
            newNum += digit * Math.pow(10, (double) reversedCount);
            reversedCount++;
        }
        return oldNum != newNum;
    }
}