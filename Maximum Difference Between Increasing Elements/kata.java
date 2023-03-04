class Solution {
    public int maximumDifference(int[] nums) {
        var min = -1;
        var maximumDifference = -1;
        for (var i = 1; i < nums.length; i++) {
            if (min == -1 || nums[i-1] < min) {
                min = nums[i-1];
            }
            if (min < nums[i]) {
                var maximumDifferenceLocal = Math.abs(nums[i]-min);
                if (maximumDifferenceLocal > maximumDifference) {
                    maximumDifference = maximumDifferenceLocal;
                }
            }
        }
        return maximumDifference;
    }
}