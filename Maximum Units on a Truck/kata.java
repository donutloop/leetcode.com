class Solution {
    public int maximumUnits(int[][] boxTypes, int truckSize) {
            Arrays.sort(boxTypes, Comparator.comparingInt(a -> a[1]));
            var sum = 0;
            for (var i = boxTypes.length-1; i >= 0; i--) {
                if (boxTypes[i][0] <= truckSize) {
                    sum += boxTypes[i][0] * boxTypes[i][1];
                    truckSize -= boxTypes[i][0];
                } else {
                    sum += truckSize * boxTypes[i][1];
                    truckSize = 0;
                }
                if (truckSize == 0) {
                    break;
                }
            }
            return sum;
    }
}