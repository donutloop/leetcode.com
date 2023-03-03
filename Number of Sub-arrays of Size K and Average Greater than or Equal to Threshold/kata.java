class Solution {
    public int numOfSubarrays(int[] arr, int k, int threshold) {
        var i = 0;
        var sum = 0;
        for (; i < k; i++) {
            sum += arr[i];
        }

        var avg = sum/k;
        var matchedCount = 0;
        if (avg >= threshold) {
            matchedCount++;
        }

        for (; i < arr.length; i++) {
            sum += arr[i];
            sum -= arr[i-k];
            avg = sum/k;
            if (avg >= threshold) {
                matchedCount++;
            }
        }

        return matchedCount;
    }
}