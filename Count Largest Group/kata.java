class Solution {
    public int countLargestGroup(int n) {
        HashMap<Integer, Integer> bucket = new HashMap<>();
        HashMap<Integer, Integer> count = new HashMap<>();
        for (var unmodifiedNum = 1; unmodifiedNum < n+1; unmodifiedNum++) {
             var sum = 0;
             var num = unmodifiedNum;

            if (num <= 9) {
                bucket.put(num, num);
                count.put(num,  1);
                continue;
            }
            do {
                var r = num % 10;
                sum =+ r;
                num = num / 10;
                if (bucket.containsKey(num)) {
                  count.put(bucket.get(num)+sum, count.getOrDefault(bucket.get(num)+sum, 0) + 1);
                  break;
                }
             } while(num != 0);

            bucket.put(unmodifiedNum, bucket.get(num)+sum);
        }

        var max = -1;
        var countMax = 0;
        for (Map.Entry<Integer, Integer> entry : count.entrySet()) {
            if (max == -1 || max < entry.getValue()) {
                max = entry.getValue();
            }
        }

        for (Map.Entry<Integer, Integer> entry : count.entrySet()) {
            if (max == entry.getValue()) {
                countMax++;
            }
        }

        return countMax;
    }
}