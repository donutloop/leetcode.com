class Solution {
    public List<List<Integer>> mergeSimilarItems(int[][] items1, int[][] items2) {
        Map<Integer, List<Integer>> mergedItems = new TreeMap<>();

        for (var item: items1) {
            mergedItems.put(item[0], new ArrayList<>(Arrays.asList(item[0], item[1])));
        }

        for (var item: items2) {
            if (mergedItems.containsKey(item[0])) {
                var element = mergedItems.get(item[0]);
                element.set(1, element.get(1)+item[1]);
                mergedItems.put(item[0], element);
            } else {
                mergedItems.put(item[0], new ArrayList<>(Arrays.asList(item[0], item[1])));
            }
        }

        var list = new ArrayList<List<Integer>>(mergedItems.values());

        return list;
    }
}