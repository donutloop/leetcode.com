class Solution {
    public int maximumXOR(int[] nums) {
        var max = 0;
        for (var num : nums) {
            max |= num;
        }
        return max;
    }
}