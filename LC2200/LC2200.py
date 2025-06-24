from typing import List


class Solution:
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        n = len(nums)
        a = [0] * (n + 1)
        print(a)
        for i in range(n):
            if nums[i] == key:
                a[max(0, i - k)] += 1
                a[min(n, i + k + 1)] -= 1
        for i in range(1, n):
            a[i] += a[i - 1]
        return [i for i in range(n) if a[i] > 0]

nums = [3,4,9,1,3,9,5]
key = 9
k = 1
sol = Solution()
print(sol.findKDistantIndices(nums, key, k))