class Solution {
    public int[] findBuildings(int[] heights) {
        var n = heights.length;
        Stack<Integer> buildings = new Stack<>();
        for (int i = 0; i < n; i++) {
            while (!buildings.empty() && heights[buildings.peek()] <= heights[i]) {
                buildings.pop();
            }
            buildings.push(i);
        }
        var result = new int[buildings.size()];
        var i = buildings.size()-1;
        while (!buildings.isEmpty()) {
            result[i] = buildings.pop();
            i--;
        }
        return result;
    }
}