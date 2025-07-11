from typing import List

class Solution:
    def countDays(self, days: int, meetings: List[List[int]]) -> int:
        meetings.sort()
        i, ans = 0, 0
        while i < len(meetings):
            j, r = i + 1, meetings[i][1]
            while j < len(meetings) and meetings[j][0] <= r:
                r = max(r, meetings[j][1])
                j += 1
            # print(f'l:{meetings[i][0]}, r:{r}')
            ans += r - meetings[i][0] + 1
            i = j
        return days - ans
    
days = int(input())
n = int(input())
meetings = [list( map(int, input().split()) ) for _ in range(n)]
ans = Solution().countDays(days, meetings)
print(ans)