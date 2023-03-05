class Solution {
    public int minimumOperations(int[] nums) {
        var countOperation = 0;
        Arrays.sort(nums);
        var sum = 0;
        for (var i = 0; i < nums.length; i++) {
            if (nums[i] == 0) {
                continue;
            }
            if (sum < nums[i]) {
                sum += Math.abs(sum-nums[i]);
                countOperation++;
            }
        }
        return countOperation;
    }
}


