import itertools

class Solution:
    def countPrefixSuffixPairs(self, words: List[str]) -> int:   
        matched = 0
        for pair in itertools.combinations(words, 2):
            if self.isPrefixAndSuffix(pair[0], pair[1]):
                matched = matched + 1

        return matched    
            
    def isPrefixAndSuffix(self, a: str, b :str) -> bool: 
        if len(a) > len(b):
            return False
        if a == b[:len(a)] and b[len(b)-len(a):] == a:
            return True
        return False         

