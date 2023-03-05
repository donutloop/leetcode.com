class Solution {
    public int maxCoins(int[] piles) {
        Arrays.sort(piles);
        var j = 0;
        var sum = 0;
        for (var i = piles.length-2; i >= j; i=i-2) {
            sum += piles[i];
            j++;
        }
        return sum;
    }
}