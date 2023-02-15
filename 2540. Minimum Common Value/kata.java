class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        var i = 0;
        var j = 0;
        HashMap<Integer, Integer> counter = new HashMap<>();
        while(i < nums1.length || j < nums2.length) {

          if (i < nums1.length) {

            var c = 0;
            if (counter.containsKey(counter.get(nums1[i]))) {
                c = counter.get(nums1[i]);
            }

            counter.put(nums1[i]+Integer.MIN_VALUE,  c + 1);
            if (counter.containsKey(nums1[i]) && counter.get(nums1[i]) > 0 && counter.get(nums1[i]+Integer.MIN_VALUE) > 0){
                  return nums1[i];
            }
          }

          if (j < nums2.length) {

            var c = 0;
            if (counter.containsKey(counter.get(nums2[j]))) {
                c = counter.get(nums2[j]);
            }

            counter.put(nums2[j], c + 1);

            if (counter.get(nums2[j]) > 0 && counter.containsKey(nums2[j]+Integer.MIN_VALUE) && counter.get(nums2[j]+Integer.MIN_VALUE) > 0){
               return nums2[j];
            }
          }

          i++;
          j++;
        }
        return -1;
    }
}