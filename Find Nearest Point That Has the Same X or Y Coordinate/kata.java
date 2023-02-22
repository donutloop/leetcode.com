class Solution {
    public int nearestValidPoint(int x, int y, int[][] points) {
        var index = -1;
        var smallestDistance = -1;
        var i = 0;
        for (var point : points) {
            if (point[0] == x || point[1] == y) {
                var newDistance = Math.abs(point[0] - x) + Math.abs(point[1] - y);
                if (index == -1) {
                    index = i;
                    smallestDistance = newDistance;
                } else if (smallestDistance > newDistance) {
                    index = i;
                    smallestDistance = newDistance;
                }
            }
            i++;
        }

        return index;
    }
}