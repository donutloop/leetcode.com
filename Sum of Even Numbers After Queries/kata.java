class Solution {
    public int[] sumEvenAfterQueries(int[] nums, int[][] queries) {
        var res = new int[queries.length];
        var sum = 0;
        for (var num : nums) {
            if (num % 2 == 0) {
                sum += num;
            }
        }
        for (var i = 0; i < queries.length; i++) {
            var oldNum = nums[queries[i][1]];
            var newNum = nums[queries[i][1]] + queries[i][0];
            if (oldNum % 2 == 0) {
                sum = sum - oldNum;
            }
            if (newNum % 2 == 0) {
                sum = sum + newNum;
            }
            nums[queries[i][1]] = newNum;
            res[i] = sum;
        }
        return res;
    }
}