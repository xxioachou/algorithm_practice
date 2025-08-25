from typing import List

class Solution:
    def minimumArea(self, grid: List[List[int]]) -> int:
        m = len(grid)    
        n = len(grid[0])
        u, d = m, -1
        l, r = n, -1
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    l, r = min(l, j), max(r, j)
                    u, d = min(u, i), max(d, i)
        if d == -1:
            return 0
        return (r - l + 1) * (d - u + 1)

m = int(input())
a = [list(map(int, input().split(','))) for _ in range(m)]
print(Solution().minimumArea(a))