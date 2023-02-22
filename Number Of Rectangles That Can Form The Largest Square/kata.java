class Solution {
    public int countGoodRectangles(int[][] rectangles) {
        var count = 0;
        var maxLen = -1;
        for (var rectangle : rectangles) {
            var minLen = min(rectangle[0], rectangle[1]);
            if (maxLen == -1 || minLen > maxLen) {
                maxLen = minLen;
                count = 1;
            } else if (minLen == maxLen) {
                count++;
            }
        }
        return count;
    }
    private int min(int a, int b) {
        if (a < b) {
            return a;
        }
        return b;
    }
}