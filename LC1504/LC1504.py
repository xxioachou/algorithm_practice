from typing import List

class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        m = len(mat)
        n = len(mat[0])
        s = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m):
            for j in range(n):
                s[i + 1][j + 1] = s[i + 1][j] + s[i][j + 1] - s[i][j] + mat[i][j]
        ans = 0
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                for x in range(1, m - i + 2):
                    l, r = 0, n - j + 1
                    while l < r:
                        mid = (l + r + 1) // 2
                        a, b = i + x - 1, j + mid - 1
                        tot = s[a][b] - s[i - 1][b] - s[a][j - 1] + s[i - 1][j - 1]
                        if tot >= x * mid:
                            l = mid
                        else:
                            r = mid - 1
                    ans += r
        return ans

rows = int(input())
a = [list(map(int, input().split(','))) for _ in range(rows)]
# print(a)
print(Solution().numSubmat(a))