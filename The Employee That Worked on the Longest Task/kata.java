class Solution {
    public int hardestWorker(int n, int[][] logs) {
        if (logs.length == 0) {
            return -1;
        }
        var id = logs[0][0];
        var taskSpanTime = logs[0][1];
        for (var i = 1; i < logs.length; i++) {
            var currentSpan = logs[i][1] - logs[i-1][1];
            if (taskSpanTime == currentSpan && id > logs[i][0]) {
                id = logs[i][0];
                taskSpanTime = currentSpan;
            } else if (taskSpanTime < currentSpan) {
                id = logs[i][0];
                taskSpanTime = currentSpan;
            }
        }
        return id;
    }
}