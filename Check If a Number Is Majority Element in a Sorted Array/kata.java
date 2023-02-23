class Solution {
    public boolean isMajorityElement(int[] nums, int target) {
         var count = 0;
         for (var num: nums) {
              if (num == target) {
                  count++;
              }
         }
         return count > nums.length/2;
    }
}