class Solution {
    public int dietPlanPerformance(int[] calories, int k, int lower, int upper) {
        var gained = 0;
        var lost = 0;

        var partialSum = 0;
        var count = 0;
        for (var i = 0; i < calories.length; i++) {
            partialSum += calories[i];
            count++;
            if (count == k) {
                if (partialSum < lower) {
                    lost++;
                } else if (partialSum > upper) {
                    gained++;
                }
                partialSum = partialSum - calories[i-k+1];
                count--;
            }
        }

        return gained-lost;
    }
}