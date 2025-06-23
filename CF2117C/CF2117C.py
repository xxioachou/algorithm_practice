import sys
from collections import deque
input = sys.stdin.readline

def main():
    t = int(input())
    for _ in range(t): 
        n = int(input())
        a = list(map(int, input().split()))
        idx = [deque() for _ in range(n + 1)]
        for i in range(n):
            idx[a[i]].append(i)

        l, r, ans = 0, 0, 1
        while r < n:
            # print(f"l: {l}, r: {r}, ans: {ans}")
            next_r = -1
            for i in range(l, r + 1):
                while len(idx[a[i]]) > 0 and idx[a[i]][0] <= r:
                    idx[a[i]].popleft()
                if len(idx[a[i]]) == 0:
                    next_r = -1
                    break
                next_r = max(next_r, idx[a[i]][0])
            if next_r == -1:
                break
            ans += 1
            l, r = r + 1, next_r
        print(ans)

def solve():
    t = int(input())
    for _ in range(t): 
        n = int(input())
        a = list(map(int, input().split()))
        st = [set() for _ in range(2)]
     
        cur, i, ans = 0, 1, 1
        st[0].add(a[0])
        while i < n:
            j = i
            st[cur ^ 1] = set()
            while j < n and len(st[cur]) > 0:
                st[cur].discard(a[j])
                st[cur ^ 1].add(a[j])
                j += 1
            if len(st[cur]) > 0:
                break

            cur ^= 1
            i, ans = j, ans + 1
        print(ans)
    
if __name__ == "__main__":
    solve()
