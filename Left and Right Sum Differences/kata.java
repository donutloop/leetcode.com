class Solution {
    public int[] leftRigthDifference(int[] nums) {
        var res = new int[nums.length];
        var j = nums.length-1;
        var leftSum = 0;
        var rightSum = 0;
        for (var i = 0; i < nums.length-1; i++) {
            leftSum += nums[i];
            rightSum += nums[j];
            if (res[i+1] > 0) {
                res[i+1] -= leftSum;
                res[i+1] = Math.abs(res[i+1]);
            } else {
                res[i+1] = leftSum;
            }
            if (res[j-1] > 0) {
                res[j-1] -= rightSum;
                res[j-1] = Math.abs(res[j-1]);
            } else {
                res[j-1] = rightSum;
            }
            j--;
        }
        return res;
    }
}