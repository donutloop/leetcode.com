class Solution {
    public int arithmeticTriplets(int[] nums, int diff) {

        Set<Integer> set = new HashSet<Integer>();
        for (var num : nums) {
            set.add(num);
        }

        var count = 0;
        for (var i = 0; i < nums.length; i++) {
            var jump = nums[i]+diff;
            if (set.contains(jump)) {
                var isTriplet = 1;
                do {
                   jump = jump+diff;
                   if (set.contains(jump)) {
                       isTriplet++;
                   } else {
                       break;
                   }

                } while(isTriplet != 2);

                if (isTriplet == 2) {
                    count++;
                }
            }
        }

        return count;
    }
}