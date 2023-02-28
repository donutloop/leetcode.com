class Solution {
    public int[] countPoints(int[][] points, int[][] queries) {

        var res = new int[queries.length];
        var i = 0;
        for (var query : queries) {
            for (var point : points) {
                // d= √(xA − xB)2 + (yA − yB)2
                var distance = Math.sqrt(Math.pow(query[0]-point[0],2) + Math.pow(query[1]-point[1],2));
                if (distance <= (double) query[2]) {
                    res[i]++;
                }
            }
            i++;
        }
        return res;
    }
}