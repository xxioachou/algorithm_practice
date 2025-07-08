
def main():
    t = int(input())
    for _ in range(t):
        n = int(input())
        d = list( map(int, input().split()) )
        seg = [list( map(int, input().split())) for _ in range(n)]
        lows, highs = [], []
        low, high = 0, 0
        lows.append(low)
        highs.append(high)

        ok = True
        for i in range(n):
            t = low
            if d[i] == -1:
                high += 1
            else:
                low, high = low + d[i], high + d[i]
            low = max(low, seg[i][0])
            high = min(high, seg[i][1])
            if low > high:
                ok = False
                break
            lows.append(low)
            highs.append(high)
        if not ok:
            print(-1)
        else:
            cur = highs[n]
            for i in range(n - 1, -1, -1):
                if d[i] != -1:
                    cur -= d[i]
                else:
                    if lows[i] <= cur and cur <= highs[i]:
                        d[i] = 0
                    else:
                        d[i] = 1
                        cur -= 1
            print(' '.join(str(i) for i in d))

            # cur = 0
            # for i in range(n):
            #     cur += d[i]
            #     assert seg[i][0] <= cur and cur <= seg[i][1]

if __name__ == '__main__':
    main()