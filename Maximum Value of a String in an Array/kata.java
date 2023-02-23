class Solution {
    public int maximumValue(String[] strs) {
        var max = 0;
        outerLoop:
        for (var str: strs) {
            var localMax = 0;
            var j = 0;
            for (var i = str.length()-1; i >= 0; i--) {
                if (Character.isLetter(str.charAt(i))) {
                    if (max < str.length()) {
                        max = str.length();
                    }
                    continue outerLoop;
                }
                var digit = Character.getNumericValue(str.charAt(i));
                localMax += (digit * Math.pow(10.0, (double) j));
                j++;
            }
            if (localMax > max) {
                max = localMax;
            }
        }
        return max;
    }
}