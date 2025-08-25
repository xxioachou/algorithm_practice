class Solution:
    def makeFancyString(self, s: str) -> str:
        ans = ''
        i = 0
        while i < len(s):
            j = i + 1
            while j < len(s) and s[i] == s[j]:
                j += 1
            ans += min(j - i, 2) * s[i]
            i = j
        return ans