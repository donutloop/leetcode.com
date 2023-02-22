class Solution {
    public int countAsterisks(String s) {
        var skip = false;
        var count = 0;
        for (var i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '|' && !skip) {
                skip = true;
                continue;
            } else if (s.charAt(i) == '|' && skip) {
                skip = false;
                continue;
            }
            if (skip) {
                continue;
            }
            if (s.charAt(i) == '*') {
                count++;
            }
        }
        return count;
    }
}