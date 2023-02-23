import java.math.BigInteger;

class Solution {
    public int countPrimeSetBits(int left, int right) {
        var count = 0;
        for (var num = left; num <= right; num++) {
            var bitCount = BigInteger.valueOf(Integer.bitCount(num));
            if (bitCount.isProbablePrime(30)) {
                count++;
            }
        }
        return count;
    }
}