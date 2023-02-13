class Solution {
    public long findTheArrayConcVal(int[] nums) {
        var j = nums.length-1;
        long c  = 0;
        var i = 0;
        for (;i < nums.length/2; i++) {
            var al = (int) (Math.log10(nums[i]) + 1);
            var bl = (int) (Math.log10(nums[j]) + 1);
            var a = 0;
            var b = 0;

            if (al == bl) {
                a = nums[i] * (int) Math.pow(10, (double)al);
                b = nums[j];
            } else {
                a = nums[i] * (int) Math.pow(10, (double)bl);
                b = nums[j];
            }
            c += a + b;
            j--;
        }

        if (nums.length % 2 != 0) {
            c += nums[i];
        }

        return c;
    }
}