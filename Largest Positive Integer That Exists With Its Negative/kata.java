import java.util.Arrays;
import java.util.Comparator;

class Solution {
    public int findMaxK(int[] nums) {

        Integer[] boxedNums = new Integer[nums.length];
        for (int i = 0; i < nums.length; i++) {
            boxedNums[i] = nums[i];
        }

        Arrays.sort(boxedNums, new Comparator<Integer>() {
            @Override
            public int compare(Integer n1, Integer n2) {
                if (n2 < 0) {
                    n2 = n2 * -1;
                }
                if (n1 < 0) {
                    n1 = n1 * -1;
                }
                return Integer.compare(n1, n2);
            }
        });

        var i = nums.length - 2;
        var k = -1;
        for (; i >= 0; i--) {
            var j = i+1;
            if (boxedNums[j].intValue() == boxedNums[i].intValue()) {
                continue;
            }

            var ni = boxedNums[i];
            if (boxedNums[i] < 0) {
                ni = ni * -1;
            }

            var nj = boxedNums[j];
            if (boxedNums[j] < 0) {
                nj = nj * -1;
            }

            if (nj.intValue() == ni.intValue()) {
                return nj;
            }
        }

        return k;
    }
}