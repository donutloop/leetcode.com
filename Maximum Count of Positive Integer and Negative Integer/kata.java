class Solution {
    public int maximumCount(int[] nums) {
        var posCount = 0;
        var negCount = 0;
        for (var num : nums) {
            if (num >= 1) {
                posCount++;
            } else if (num <= -1) {
                negCount++;
            }
        }
        if (negCount > posCount) {
            return negCount;
        }
        return posCount;
    }
}