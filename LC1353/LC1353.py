import heapq
from typing import List

class Solution:
    def maxEvents(self, events: List[List[int]]) -> int:
        LEN = 100000

        l = [list() for _ in range(LEN + 1)]
        for i in range(len(events)):
            l[events[i][0]].append(i)
        h = []
        ans = 0
        for i in range(1, LEN + 1):
            for idx in l[i]:
                heapq.heappush(h, (events[idx][1], idx))
            while len(h) > 0 and h[0][0] < i:
                heapq.heappop(h)
            if len(h) > 0:
                # t = heapq.heappop(h)
                # print(t)
                heapq.heappop(h)
                ans += 1
        return ans

        
n = int(input())
e = [list(map(int, input().split())) for _ in range(n)]
print(e)
print(Solution().maxEvents(e))
# print(Solution().maxEvents([[1,2],[2,3],[3,4],[1,2]]))
# print(Solution().maxEvents([[1,2],[2,3],[3,4]]))