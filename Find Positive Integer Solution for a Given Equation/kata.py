package kata

"""
   This is the custom function interface.
   You should not implement it, or speculate about its implementation
   class CustomFunction:
       # Returns f(x, y) for any given positive integers x and y.
       # Note that f(x, y) is increasing with respect to both x and y.
       # i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
       def f(self, x, y):
  
"""

import itertools

class Solution:
    def findSolution(self, customfunction: 'CustomFunction', z: int) -> List[List[int]]:
        
        numbersPairs = []
        numbers = []
        for i in range(1, z+1):
            numbers.append(i)
            numbersPairs.append([i, i])
            
        numbersPairs.extend(list(itertools.permutations(numbers, 2)))
        
        newNumbersPairs = []
        for i in range(len(numbersPairs)):
            z0 = customfunction.f(numbersPairs[i][0], numbersPairs[i][1])
            if z == z0:
                newNumbersPairs.append(numbersPairs[i]) 
                
       
        return newNumbersPairs
            
            
        
