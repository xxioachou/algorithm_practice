from typing import List
from collections import deque

class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        startTime.append(eventTime)
        endTime.append(eventTime)

        n = len(startTime)
        q = []
        q.append(0)
        r, sum, j = 0, 0, 1
        ans = 0
        for i in range(n):
            # print(f"l: {startTime[i]}, r: {endTime[i]}")
            sum += startTime[i] - r
            q.append(startTime[i] - r)
            r = endTime[i]
            while len(q) - j > k + 1:
                sum -= q[j]
                j += 1
            if len(q) - j == k + 1:
                ans = max(ans, sum)
        return ans

et = int(input())
k = int(input())
st = list( map(int, input().split()) )
ed = list( map(int, input().split()) )
print(Solution().maxFreeTime(et, k, st, ed))