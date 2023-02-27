class Solution {
    public int findTheWinner(int n, int k) {
        Queue<Integer> queue = new LinkedList<>();
        for (var i = 1; i < n+1; i++) queue.add(i);
        var i = 0;
        while(queue.size() != 1) {
            var element = queue.poll();
            i++;
            if (i == k) {
                i = 0;
            } else {
                queue.add(element);
            }
        }
        return queue.poll();
    }
}