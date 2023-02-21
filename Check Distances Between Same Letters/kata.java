class Solution {
    public boolean checkDistances(String s, int[] distance) {
        for (var i = 0; i < s.length(); i++) {
            var distanceOfChar = distance[s.charAt(i)- 97];
            var rightPosNewChar = distanceOfChar + i + 1;
            var leftPosNewChar =  i - distanceOfChar - 1;
            var left = true;
            var right = true;
            if (rightPosNewChar < s.length()) {
                if (s.charAt(i) == s.charAt(rightPosNewChar)) {
                    right = false;
                }
            }
            if (leftPosNewChar >=  0) {
                if (s.charAt(i) == s.charAt(leftPosNewChar)) {
                    left = false;
                }
            }
            if (left && right) {
                return false;
            }
        }
        return true;
    }
}