class Solution {
    public int maxDepth(String s) {
        var openParentheses = 0;
        var maxParentheses = 0;
        for (var i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                openParentheses++;
            } else if (s.charAt(i) == ')') {
                if (maxParentheses < openParentheses) {
                    maxParentheses =  openParentheses;
                }
                openParentheses--;
            } else if (Character.isDigit(s.charAt(i)) || s.charAt(i) == '+' || s.charAt(i) == '-' || s.charAt(i) == '/' || s.charAt(i) == '*') {
                if (maxParentheses < openParentheses) {
                    maxParentheses =  openParentheses;
                }
            }
        }
        return maxParentheses;
    }
}