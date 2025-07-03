import math

class Solution:
    def dfs(self, k: int) -> int:
        print(k)
        if k == 1:
            return 0
        p = math.ceil(math.log2(k))
        seg = 2**p // 2
        return (self.dfs(k - seg) + 1) % 26
    def kthCharacter(self, k: int) -> str:
        return chr(self.dfs(k) + 97)
    
k = int(input())
print(Solution().kthCharacter(k))
# print(Solution().kthCharacter(10))
# print(math.ceil(math.log2(k)))