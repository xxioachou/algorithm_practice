from typing import List

class Solution:
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        ans = []
        m = len(mat)
        n = len(mat[0])
        def dfs(x:int, y:int, dir:int):
            if len(ans) >= m * n:
                return
            if x < m and y < n:
                ans.append(mat[x][y])
            if dir == 0:
                if x != 0 and y != n - 1:
                    dfs(x - 1, y + 1, dir)
                else:
                    dfs(x, y + 1, dir ^ 1)
            else:
                if x != m - 1 and y != 0:
                    dfs(x + 1, y - 1, dir)
                else:
                    dfs(x + 1, y, dir ^ 1)
        dfs(0, 0, 0)
        return ans
    
m = int(input())
a = [list(map(int, input().split(','))) for _ in range(m)]
res = Solution().findDiagonalOrder(a)
print('[' + ','.join([str(x) for x in res]) + ']')