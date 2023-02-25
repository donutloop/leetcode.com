class Solution {
    public int minStartValue(int[] nums) {
        var n = nums.length+1;
        var sums = new int[n];
        Integer min = null;

        for (var k = 1; k < n; k++) {
            sums[k] = sums[k-1] + nums[k-1];
        }

        for (var i = 1; i < sums.length; i++) {
            if (sums[i] < 1 && (min == null || sums[i] < min)) {
                min = sums[i];
            }
        }

        if (min == null) {
            return 1;
        }

        sums[0] = min.intValue();
        if (sums[0] < 0)
            sums[0] = sums[0] * -1;

        for (var k = 1; k < n; k++) {
            sums[k] = sums[k-1] + nums[k-1];
        }

        var correction = 0;
        min = null;
        for (var num: sums) {
            if (num < 1 && (min == null || num < min)) {
                min = num;
            }
        }

        if (min != null) {
            correction = Math.abs(min.intValue())+1;
        }

        return sums[0]+correction;
    }
}