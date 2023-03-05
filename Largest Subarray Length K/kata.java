class Solution {
    public int[] largestSubarray(int[] nums, int k) {
        var max = 0l;
        var res = new int[k];
        var index = 0;
        for (var i = 0; i < nums.length-k+1; i++) {
            var localMax = nums[i];
            if (localMax > max) {
                max = localMax;
                index = i;
            }
        }

        var j = 0;
        for (var i = index; i < index+k; i++) {
            res[j] = nums[i];
            j++;
        }

        return res;
    }
}