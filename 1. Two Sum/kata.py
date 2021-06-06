class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        from itertools import permutations
        perm = permutations(nums, 2)
        a = None
        b = None
        for pair in perm:
            if target == pair[0] + pair[1]:
                a = pair[0]
                b = pair[1]
        j = 0
        k = 0
        for i in range(len(nums)):
            if nums[i] == b:
                j = i
            if nums[i] == a:
                k = i

        return [j, k]