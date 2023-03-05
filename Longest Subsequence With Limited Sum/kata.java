class Solution {
    public int[] answerQueries(int[] nums, int[] queries) {
        var res = new int[queries.length];
        var i = 0;
        Arrays.sort(nums);
        for (var  query : queries) {
            var sum = 0;
            var subsequenceCount = 0;
            for (var num : nums) {
                if (sum + num > query) continue;
                sum += num;
                subsequenceCount++;
            }
            res[i] = subsequenceCount;
            i++;
        }
        return res;
    }
}