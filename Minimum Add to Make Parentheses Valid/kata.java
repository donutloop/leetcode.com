class Solution {
    public int minAddToMakeValid(String s) {
        var count = 0;
        var stack = new Stack<Character>();
        for (var i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') stack.add('(');
            if (s.charAt(i) == ')') {
                if (stack.isEmpty()) {
                    count++;
                } else {
                    stack.pop();
                }
            }

        }
        return count + stack.size();
    }
}