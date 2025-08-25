from typing import List

class Solution:
    def qpow(self, a, b, mod):
        ret = 1
        while b > 0:
            if b & 1:
                ret = ret * a % mod;
            b >>= 1
            a = a * a % mod
        return ret % mod
    
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        powers = [0]
        for i in range(30):
            if n >> i & 1:
                powers.append(1 << i)
        # print(powers)
        
        MOD = 1000000007
        pre_powers = [1]
        for i in range(1, len(powers), 1):
            pre_powers.append(pre_powers[i - 1] * powers[i] % MOD)

        return [(pre_powers[queries[i][1] + 1] * self.qpow(pre_powers[queries[i][0]], MOD - 2, MOD) % MOD) for i in range(len(queries))]
