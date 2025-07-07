from typing import List
from math import ceil, log2

'''
x

xx
xxxx
xxxxxxxx
xxxxxxxxxxxxxxxx
......

final len: 1 * 2^len(operations)

def dfs(k):
1. k == 1, return 0
2. k != 1
    k <= 2^x
    x = ceil(log2(k))
    p = 2 ** (x - 1)
    return (dfs(k - p) + operation[x - 1]) % 26
'''
class Solution:
    def __init__(self):
        self.operations = []

    def kthCharacter(self, k: int, operations: List[int]) -> str:
        self.operations = operations
        return chr(self.dfs(k) + 97)

    def dfs(self, k: int) -> int:
        if k == 1:
            return 0
        
        x = ceil(log2(k))
        assert x > 0
        p = 2 ** (x - 1)
        return (self.dfs(k - p) + self.operations[x - 1]) % 26        

k = int(input())
op = list(map(int, input().split()))
print(k, op, len(op), 2 ** len(op))
print(Solution().kthCharacter(k, op))