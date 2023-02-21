class Solution {
    public int numberOfCuts(int n) {
        if (n == 1) {
            return 0;
        }
        if (n % 2 == 1) {
            return n;
        }
        return n / 2;
    }
}