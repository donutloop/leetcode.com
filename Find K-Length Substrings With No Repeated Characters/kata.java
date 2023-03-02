class Solution {
    public int numKLenSubstrNoRepeats(String s, int k) {
        if (s.length() < k) {
            return 0;
        }

        var characterCount = 0;
        var wordCount = 0;
        HashMap<Character, Integer> windowTracker = new HashMap<>();
        for (var i = 0; i < s.length(); i++) {
            if (windowTracker.containsKey(s.charAt(i))) {
                var otherIndex = windowTracker.get(s.charAt(i));
                i = otherIndex + 1;
                windowTracker.clear();
                characterCount = 0;
            }
            windowTracker.put(s.charAt(i), i);
            characterCount++;
            if (characterCount == k) {
                wordCount++;
                windowTracker.remove(s.charAt(i - (k-1)));
                characterCount = k-1 ;
            }
        }

        return wordCount;
    }
}