class Solution {
    public List<Integer> minSubsequence(int[] nums) {
        Arrays.sort(nums);
        var maximumTotalSum = 0;
        for (var num:  nums) {
            maximumTotalSum += num;
        }

        var partialSum = 0;
        List<Integer> foundNumbers = new ArrayList<>();
        for (var i = nums.length-1; i >= 0; i--) {
            partialSum += nums[i];
            foundNumbers.add(nums[i]);
            maximumTotalSum -= nums[i];
            if (partialSum > maximumTotalSum) {
                break;
            }
        }

        return foundNumbers;
    }
}