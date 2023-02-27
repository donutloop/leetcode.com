class Solution {
    public int minPartitions(String n) {
        if (n.length() == 1) {
            return n.charAt(0)-48;
        }
        var lastDigit = -1;
        for (var i = 0; i < n.length(); i++) {
            var a = n.charAt(i)-48;
            if (lastDigit == -1 || lastDigit < a) {
                lastDigit = a;
            }
            if (lastDigit == 9) {
                break;
            }
        }
        return lastDigit;
    }
}