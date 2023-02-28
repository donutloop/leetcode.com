class Solution {
    public int maxWidthOfVerticalArea(int[][] points) {
        Integer[][] convertedPoints = new Integer[points.length][points[0].length];
        for (int i = 0; i < points.length; i++) {
            for (int j = 0; j < points[i].length; j++) {
                convertedPoints[i][j] = Integer.valueOf(points[i][j]);
            }
        }

        Arrays.sort(convertedPoints, new Comparator<Integer[]>() {
            public int compare(Integer[] p1, Integer[] p2) {
                return p1[0].compareTo(p2[0]);
            }
        });

        var max = 0;
        for (int i = 1; i < convertedPoints.length; i++) {
            var d = convertedPoints[i][0]-convertedPoints[i-1][0];
            if (d > max) {
                max = d;
            }
        }
        return max;
    }
}