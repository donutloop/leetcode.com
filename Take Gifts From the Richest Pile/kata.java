class Solution {
    public long pickGifts(int[] gifts, int k) {

        PriorityQueue<Integer> q = new PriorityQueue<Integer>(gifts.length, Collections.reverseOrder());
        for (var gift: gifts) {
            q.add(gift);
        }

        while(k >= 1) {
             var gift = q.poll();
             q.add((int) Math.floor(Math.sqrt((double) gift)));
             k--;
        }

        var sum = 0l;
        while(!q.isEmpty()) {
            sum += q.poll().intValue();
        }

        return sum;
    }
}