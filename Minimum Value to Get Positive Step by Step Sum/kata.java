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

        min = null;
        for (var k = 1; k < n; k++) {
            sums[k] = sums[k-1] + nums[k-1];
            if (sums[k] < 1 && (min == null || sums[k] < min)) {
                min = sums[k];
            }
        }

        var correction = 0;
        if (min != null) {
            correction = Math.abs(min.intValue())+1;
        }

        return sums[0]+correction;
    }
}