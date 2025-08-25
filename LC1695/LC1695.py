from typing import List

class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        V = 10000
        j, n = 0, len(nums)       
        mp = [0] * (V + 1)
        sum, ans = 0, 0
        for i in range(n):
            while j < n and mp[nums[j]] < 1:
                mp[nums[j]] += 1
                sum += nums[j]
                j += 1
            ans = max(ans, sum)
            mp[nums[i]] -= 1
            sum -= nums[i]
        return ans

a = list(map(int, input().split(',')))
print(a)
print(Solution().maximumUniqueSubarray(a))