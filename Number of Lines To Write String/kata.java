class Solution {
    public int[] numberOfLines(int[] widths, String s) {
        var res = new int[2];

        final var threshold = 100;
        var currentLineWidth = 0;
        for (var i = 0; i < s.length(); i++) {
            if (currentLineWidth + widths[s.codePointAt(i)-97] > threshold) {
                currentLineWidth = widths[s.codePointAt(i)-97];
                res[0]++;
            } else {
                currentLineWidth = currentLineWidth + widths[s.codePointAt(i)-97];
            }
        }

        if (currentLineWidth > 0) {
            res[1] = currentLineWidth;
            res[0]++;
        }

        return res;
    }
}