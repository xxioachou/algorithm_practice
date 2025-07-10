from typing import List

'''
    a[i] = startTime[i + 1] - endTime[i]
    length[i] = endTime[i] - startTime[i]
    pre[i] = max(a[1 : i+1])
    suf[i] = max(a[i : n + 2])
    ans = max(a[i] + a[i + 1])
    if pre[i - 1] >= length[i] or
        suf[i + 2] >= length[i]
        ans = max(ans, a[i] + a[i + 1] + length[i])
'''
class Solution:
    def maxFreeTime(self, eventTime: int, startTime: List[int], endTime: List[int]) -> int:
        n = len(startTime)
        a = [0]
        length = [0]
        pre = 0
        for i in range(n):
            a.append(startTime[i] - pre)
            length.append(endTime[i] - startTime[i])
            pre = endTime[i]
        a.append(eventTime - pre)

        # print(a)
        # print(length)

        suf = [0] * (n + 3)
        for i in range(n + 1, 0, -1):
            suf[i] = max(suf[i + 1], a[i])
         
        ans, pre = 0, 0
        for i in range(1, n + 1):
            ans = max(ans, a[i] + a[i + 1])
            if pre >= length[i] or suf[i + 2] >= length[i]:
                ans = max(ans, a[i] + a[i + 1] + length[i])
            pre = max(pre, a[i])
        return ans

eve = int(input())
st = list( map(int, input().split()) )
ed = list( map(int, input().split()) )
print(Solution().maxFreeTime(eve, st, ed))
