import math
        
class Solution(object):
    def countSubstrings(self, s, c):
        """
        :type s: str
        :type c: str
        :rtype: int
        """
        count = 0
        for char in s:
            if char == c:
                count = count + 1

        return count + math.comb(count,2)
        