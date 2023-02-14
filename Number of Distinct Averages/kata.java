import java.util.Arrays;
import java.util.HashSet;

class Solution {
    public int distinctAverages(int[] nums) {
        Arrays.sort(nums);
        var i = 0;
        var j = nums.length-1;
        var distinctAverages = new HashSet<Double>();
        for (; i < nums.length; i++) {
            distinctAverages.add(((double) (nums[i] + nums[j])) / 2);
            j--;
        }
        return distinctAverages.size();
    }
}