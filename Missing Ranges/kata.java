class Solution {
    public List<String> findMissingRanges(int[] nums, int lower, int upper) {
        var ranges = new ArrayList<String>();
        if (nums.length == 0) {
            if (lower != upper) {
                ranges.add( Integer.toString(lower) + "->" + Integer.toString(upper));
            } else {
                ranges.add( Integer.toString(lower));
            }
            return ranges;
        }

        if (nums[0] != lower) {
            if (nums[0]-1 == lower) {
                ranges.add(Integer.toString(lower));
            } else {
                ranges.add(Integer.toString(lower) + "->" + Integer.toString(nums[0]-1));
            }
        }

        for (var i = 1; i < nums.length; i++) {
            var d = nums[i] - nums[i-1];
            if (d == 1) {
               continue;
            }
            var numA = nums[i-1] + 1;
            var numB = nums[i] - 1;
            if (numA != numB) {
                ranges.add( Integer.toString(numA) + "->" + Integer.toString(numB));
            } else {
                ranges.add( Integer.toString(numA));
            }
        }

        if (nums[nums.length-1] != upper) {
            if (nums[nums.length-1]+1 == upper) {
                ranges.add( Integer.toString(upper));
            } else {
                ranges.add(Integer.toString(nums[nums.length-1]+1) + "->" + Integer.toString(upper));
            }
        }

        return ranges;
    }
}