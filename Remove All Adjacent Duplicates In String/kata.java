class Solution {
    public String removeDuplicates(String s) {
        var sb = new StringBuilder(s);
        for (var i = 1; i < sb.length(); i++) {
            if (sb.charAt(i) == sb.charAt(i-1)) {
                sb.deleteCharAt(i);
                sb.deleteCharAt(i-1);
                if (i - 2 >= 0) i = i - 2;
                else i = 0;
            }
        }
        return sb.toString();
    }
}