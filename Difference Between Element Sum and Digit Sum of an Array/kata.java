class Solution {
    public int differenceOfSum(int[] nums) {
        var elementSum = 0;
        var digitSum = 0;
        for (var num : nums) {
            elementSum += num;
            while (num != 0) {
                var digit = num % 10;
                num = num / 10;
                digitSum += digit;
            }

        }

        return (int) Math.abs((double) elementSum-digitSum);
    }
}