from typing import List

class Solution:
    def zeroFilledSubarray(self, nums: List[int]) -> int:
        ans = 0
        i, j = 0, 0
        while i < len(nums):
            if nums[i] != 0:
                i += 1
                continue
            j = i + 1
            while j < len(nums) and nums[j] == 0:
                j += 1
            x = j - i
            ans += (x + 1) * x // 2
            i = j
        return ans