
def main():
    t = int(input())
    for _ in range(t):
        n = int(input())
        a = list(map(int, input().split()))
        cnt = [[0, 0] for _ in range(30)]
        for i in range(len(cnt)):
            for x in a:
                cnt[i][x >> i & 1] += 1

        ans = 0
        for x in a:
            res = 0
            for i in range(30):
                res += (1 << i) * cnt[i][(x >> i & 1) ^ 1]
            ans = max(ans, res)
        print(ans)
    
if __name__ == "__main__":
    main()
