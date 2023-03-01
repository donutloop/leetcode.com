class Solution {
    public int arithmeticTriplets(int[] nums, int diff) {

        Set<Integer> set = new HashSet<Integer>();
        for (var num : nums) {
            set.add(num);
        }

        var tripletCount = 0;
        for (var i = 0; i < nums.length; i++) {
            var jump = nums[i]+diff;
            if (set.contains(jump) && set.contains(jump+diff)) {
                tripletCount++;
            }
        }

        return tripletCount;
    }
}