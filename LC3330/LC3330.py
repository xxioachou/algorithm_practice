class Solution:
    def possibleStringCount(self, word: str) -> int:
        i = 0
        ans = 1
        while i < len(word):
            j = i + 1
            while j < len(word) and word[j] == word[j - 1]:
                j += 1
            ans += (j - i - 1)
            i = j
        return ans