class Solution {
    public boolean canAttendMeetings(int[][] intervals) {
        // sort intervals by first and second element in []int asscending order
        Arrays.sort(intervals, new Comparator<int[]>() {
            @Override
            public int compare(int[] a, int[] b) {
                // Compare the first element of each sub-array
                var compared = Integer.compare(a[0], b[0]);
                if(compared > 0) {
                    // a is greater than b
                    return compared;
                } else if(compared < 0) {
                    // a is smaller than b
                    return compared;
                }
                // if first element of a and b is equal then compare the second element of each sub array
                return Integer.compare(a[1], b[1]);
            }
        });

        for (int i = 1; i < intervals.length; i++) {
           if (intervals[i-1][1] <= intervals[i][0]) {
                // if end of interval of a is smaller or equal to start of interval b then no overlapp
               continue;
           }
           return false;
        }

        return true;
    }
}