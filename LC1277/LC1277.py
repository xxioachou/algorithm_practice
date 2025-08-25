from typing import List

class Solution:
    def countSquares(self, matrix: List[List[int]]) -> int:
        m = len(matrix)
        n = len(matrix[0])
        s = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m):
            for j in range(n):
                s[i + 1][j + 1] = s[i][j + 1] + s[i + 1][j] - s[i][j] + matrix[i][j]
        ans = 0
        for x in range(1, min(m, n) + 1):
            for i in range(1, m - x + 2):
                for j in range(1, n - x + 2):
                    a, b = i + x - 1, j + x - 1
                    sum = s[a][b] - s[i - 1][b] - s[a][j - 1] + s[i - 1][j - 1]
                    if sum == x * x:
                        ans += 1
        return ans

m = int(input()) 
a = [list(map(int, input().split(','))) for _ in range(m)]
print(Solution().countSquares(a))