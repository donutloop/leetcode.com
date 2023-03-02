class Solution {
    public int[] distinctNumbers(int[] nums, int k) {
        var ans = new int[nums.length-k+1];
        var count = 0;
        var j = 0;
        Map<Integer, Integer> unique = new HashMap<>();
        for (var i = 0; i < nums.length; i++) {
            unique.put(nums[i], unique.getOrDefault(nums[i], 0)+1);
            count++;
            if (count == k) {
                ans[j] = unique.size();
                var n = nums[i-(k-1)];
                var c  = unique.get(n);
                if (c > 1) {
                    c--;
                    unique.put(n, c);
                } else if (c == 1) {
                    unique.remove(n);
                }
                count--;
                j++;
            }
        }
        return ans;
    }
}