class Solution {
    public int countDigits(int num) {
        var c = 0;
        var oldNum = num;
        do {
         var r = num % 10;
         num = num / 10;
         if (oldNum % r == 0) {
             c++;
         }
        } while (num != 0);

        return c;
    }
}