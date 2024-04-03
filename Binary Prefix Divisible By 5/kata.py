class Solution(object):
    def prefixesDivBy5(self, nums):
        """
        :type nums: List[int]
        :rtype: List[bool]
        """
        result = [False] * len(nums)  
        tracker = 1  
        i = 0
        for i, num in enumerate(nums):
            if num != 0:  
                break
            else:
                result[i] = True  

        i += 1  
        for num in nums[i:]:
            if num == 0:
                tracker = tracker << 1  
                if tracker % 5 == 0:
                    result[i] = True
            else:
                tracker = tracker << 1  
                tracker = tracker | 1  
                if tracker % 5 == 0:
                    result[i] = True
            i += 1

        return result
        