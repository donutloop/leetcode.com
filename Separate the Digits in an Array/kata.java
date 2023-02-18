class Solution {
    public int[] separateDigits(int[] nums) {
        ArrayList<Integer> newNums = new ArrayList<Integer>();
        for (var num : nums) {
            var digitCount = ((int) Math.log10((double) num));
            for (; digitCount >= 0; digitCount--) {
              var digit = (num/(int) Math.pow(10, (double) digitCount)) % 10;
              newNums.add(digit);
            }
        }
        var res = new int[newNums.size()];
        for (var i = 0; i < newNums.size(); i++) {
            res[i] = newNums.get(i);
        }
        return res;
    }
}