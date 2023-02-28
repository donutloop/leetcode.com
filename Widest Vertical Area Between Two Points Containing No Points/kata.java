class Solution {
    public int maxWidthOfVerticalArea(int[][] points) {
        Integer[] convertedPoints = new Integer[points.length];
        for (int i = 0; i < points.length; i++) {
             convertedPoints[i] = Integer.valueOf(points[i][0]);
        }

        Arrays.sort(convertedPoints, new Comparator<Integer>() {
            public int compare(Integer p1, Integer p2) {
                return p1.compareTo(p2);
            }
        });

        var max = 0;
        for (int i = 1; i < convertedPoints.length; i++) {
            var d = convertedPoints[i]-convertedPoints[i-1];
            if (d > max) {
                max = d;
            }
        }
        return max;
    }
}