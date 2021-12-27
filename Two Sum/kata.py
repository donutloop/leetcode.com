package kata

from itertools import permutations  
        

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        perm = permutations(nums, 2)
        a = None
        b = None
        for pair in perm:
            if target == pair[0] + pair[1]:
                a = pair[0]
                b = pair[1]  
                break
        
        j = None
        k = None
        for i in range(len(nums)):
            if k != None and j != None:
                break
            
            if j == None and nums[i] == b:
                j = i
            elif k == None and nums[i] == a:
                k = i    
        
        return [j, k]
