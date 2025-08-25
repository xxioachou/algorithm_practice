class Solution:
    def isValid(self, word: str) -> bool:
        if len(word) < 3:
            return False
        c, v = 0, 0
        for ch in word:
            if ch.isdigit():
                continue
            elif ch in 'aeiou' or ch in 'AEIOU':
                v += 1
            elif ch.isalpha():
                c += 1
            else:
                return False
        return c > 0 and v > 0
    
s = input()
print(Solution().isValid(s))