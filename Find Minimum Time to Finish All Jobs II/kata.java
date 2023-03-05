class Solution {
    public int minimumTime(int[] jobs, int[] workers) {
        Arrays.sort(jobs);
        Arrays.sort(workers);
        var minAmount = -1d;
        for (var i = workers.length-1; i >= 0; i--) {
            var amount =  Math.ceil((double) jobs[i] / (double) workers[i]);
            if (minAmount == -1 || amount > minAmount) {
                minAmount =  amount;
            }
        }
        return (int) minAmount;
    }
}