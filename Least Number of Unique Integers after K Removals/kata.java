class Solution {
    public int findLeastNumOfUniqueInts(int[] arr, int k) {

        Map<Integer, Integer> counter = new LinkedHashMap<>();
        for (var num: arr) {
            counter.put(num, counter.getOrDefault(num, 0)+1);
        }

        PriorityQueue<Integer> pq = new PriorityQueue<>();
        for (Map.Entry<Integer, Integer> entry : counter.entrySet()) {
            pq.add(entry.getValue());
        }

        while(k != 0) {
            var value = pq.poll();
            value--;
            if (value > 0) pq.add(value);
            k--;
        }

        return pq.size();
    }
}