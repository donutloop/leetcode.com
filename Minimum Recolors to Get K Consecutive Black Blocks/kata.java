class Solution {
    public int minimumRecolors(String blocks, int k) {
        var i = 0;
        var whiteColorCount = 0;
        for ( ;i < k; i++) {
            if (blocks.charAt(i) == 'W') {
                whiteColorCount++;
            }
        }
        var min = whiteColorCount;
        for (i = k; i < blocks.length(); i++) {
            if (blocks.charAt(i) == 'W') {
                whiteColorCount++;
            }
            if (blocks.charAt(i-k) == 'W') {
                whiteColorCount--;
            }
            if (min > whiteColorCount) {
                min = whiteColorCount;
            }
        }

        return min;
    }
}