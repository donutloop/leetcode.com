import java.util.Arrays;

class Solution {
    public int maxNumberOfApples(int[] weight) {
        Arrays.sort(weight);
        var total = 0;
        var count = 0;
        for (var item: weight) {
            if (total + item > 5000) {
                break;
            }
            total = total + item;
            count++;
        }
        return count;
    }
}