class Solution {
    public int findLeastNumOfUniqueInts(int[] arr, int k) {

        Map<Integer, Integer> counter = new LinkedHashMap<>();
        for (var num: arr) {
            counter.put(num, counter.getOrDefault(num, 0)+1);
        }

        PriorityQueue<Integer> pq = new PriorityQueue<>();
        counter.forEach((key, value) -> pq.add(value));

        while(k != 0) {
            var value = pq.poll();
            value--;
            if (value > 0) pq.add(value);
            k--;
        }

        return pq.size();
    }
}