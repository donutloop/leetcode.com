class Solution {
    public int minimumDeletions(int[] nums) {
        if (nums.length == 2) {
            return 2;
        }
        if (nums.length == 1) {
            return 1;
        }

        Integer min = null;
        var indexMax = 0;
        Integer max = null;
        var indexMin = 0;
        for (var i = 0; i < nums.length; i++) {
            if (max == null || max < nums[i]) {
                max = nums[i];
                indexMax = i;
            }
            if (min == null || min > nums[i]) {
                min = nums[i];
                indexMin = i;
            }
        }

        var a = Math.abs(0-(indexMax+1));
        var b = Math.abs(nums.length-indexMin);
        var c = Math.abs(0-(indexMin+1));
        var d = Math.abs(nums.length-indexMax);

        var p = 0;
        if (a == c ) {
            p = a;
        } else if (c > a) {
            p = c;
        } else {
            p = a;
        }

        var k = 0;
        if (b == d ) {
            k = d;
        } else if (b > d) {
            k = b;
        } else {
            k = d;
        }

        if (k <= p && k <= a+b && k <= c+d) {
            return k;
        }

        if (p <= k && p <= a+b && p <= c+d) {
            return p;
        }
        if (c+d <= k && c+d <= a+b && c+d <= p) {
            return c+d;
        }

        if (a+b <= k && a+b <= c+d && a+b <= p) {
            return a+b;
        }

        return a+b;
    }
}
