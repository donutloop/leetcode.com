class Solution {
    public int averageValue(int[] nums) {
        var sum = 0;
        var count = 0;

        for (var i = 0; i < nums.length; i++) {
             if (nums[i] % 2 == 0 && nums[i] % 3 == 0) {
                 sum += nums[i];
                 count++;
             }
        }

        if (sum == 0) {
            return 0;
        }

        return (int) Math.floor(sum / count);
    }
}