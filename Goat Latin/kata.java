class Solution {
    public String toGoatLatin(String sentence) {
        var sb = new StringBuilder("");
        var aSuffix = new StringBuilder("a");
        final var maSuffix = "ma";
        var firstChar = ' ';
        for (var i = 0; i < sentence.length(); i++) {
            if (i == 0 || (i > 0 && sentence.charAt(i-1) == ' ')) {
                if (sentence.charAt(i) != 'a'
                    && sentence.charAt(i) != 'e'
                    && sentence.charAt(i) != 'i'
                    && sentence.charAt(i) != 'o'
                    && sentence.charAt(i) != 'u'
                    && sentence.charAt(i) != 'A'
                    && sentence.charAt(i) != 'E'
                    && sentence.charAt(i) != 'I'
                    && sentence.charAt(i) != 'O'
                    && sentence.charAt(i) != 'U') {
                    firstChar = sentence.charAt(i);
                    if (sentence.length()-1 == i) {
                        sb.append(firstChar);
                        sb.append(maSuffix);
                        sb.append(aSuffix);
                    }
                } else {
                    sb.append(sentence.charAt(i));
                    if (sentence.length()-1 == i) {
                        sb.append(maSuffix);
                        sb.append(aSuffix);
                    }
                }
            } else if (sentence.charAt(i) == ' ' || sentence.length()-1 == i) {
              if (sentence.length()-1 == i) {
                 sb.append(sentence.charAt(i));
              }

              if (firstChar != ' ') {
                  sb.append(firstChar);
                  firstChar = ' ';
              }
              sb.append(maSuffix);
              sb.append(aSuffix);
              aSuffix.append('a');

              if (sentence.length()-1 != i) {
                 sb.append(' ');
              }
            } else {
                sb.append(sentence.charAt(i));
            }
        }
        return sb.toString();
    }
}