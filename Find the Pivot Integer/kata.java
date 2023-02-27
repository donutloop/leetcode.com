class Solution {
    public int pivotInteger(int n) {
        var sum = 0;
        for (var num = 1; num < n+1; num++) {
            sum += num;
            var count = n - num;
            if (sum == (count+1) * num + (count*(count+1)/2)) {
                return num;
            }
        }
        return -1;
    }
}