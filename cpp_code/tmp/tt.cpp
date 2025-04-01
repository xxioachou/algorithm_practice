#include <bits/stdc++.h>
using namespace std;

const int N = 7;
const int M = 1 << N;
const double INF = 1e15;

int n, x, y[N], m[N];
int a[N][N], b[N][N];
double f[N][N][M];
bool vis[N][N][M];
double add, ans;

double dis(int x1, int y1, int x2, int y2) {
    double dx = x1 - x2;
    double dy = y1 - y2;
    return sqrtl(dx * dx + dy * dy);
}

void init() {
    for (int i = 0; i < n; i ++) {
        for (int j = 0; j < m[i]; j ++) {
            for (int k = 0; k < (1 << n); k ++) {
                f[i][j][k] = INF;
                vis[i][j][k] = false;
            }
        }
    }
}

double dfs(int x, int y, int st) {
    if (vis[x][y][st] || (st == (1 << x))) {
        return f[x][y][st];
    }
    f[x][y][st] = INF;
    for (int i = 0; i < n; i ++) {
        if ((st >> i & 1) == 0 || (i == x)) {
            continue;
        }

        for (int j = 0; j < m[i]; j ++) {
            f[x][y][st] = min(f[x][y][st], dfs(i, j, st ^ (1 << x)) + dis(a[i][j], b[i][j], a[x][y], b[x][y]));
        }
    }
    vis[x][y][st] = true;
    return f[x][y][st];
}

void calc(int x, int y) {
    for (int i = 0; i < n; i ++) {
        for (int j = 0; j < m[i]; j ++) {
            if (i == x && j != y) {
                continue;
            }
            if (dfs(i, j, (1 << n) - 1) + dis(a[i][j], b[i][j], a[x][y], b[x][y]) < ans) {
                ans = dfs(i, j, (1 << n) - 1) + dis(a[i][j], b[i][j], a[x][y], b[x][y]);
            }
            ans = min(ans, dfs(i, j, (1 << n) - 1) + dis(a[i][j], b[i][j], a[x][y], b[x][y]));
        }
    }
}



void solve() {
    // 当前在第 i 个图形上的第 j 个点上，走过的图形状态为 k 的最小距离
    cin >> n >> x;
    for (int i = 0; i < n; i ++) {
        cin >> y[i];
    }
    for (int i = 0; i < n; i ++) {
        cin >> m[i];
        for (int j = 0; j < m[i]; j ++) {
            cin >> a[i][j] >> b[i][j];
        }
        double res = 0;
        for (int j = 0; j < m[i]; j ++) {
            res += dis(a[i][j], b[i][j], a[i][(j + 1) % m[i]], b[i][(j + 1) % m[i]]);
        }
        add += res / y[i];
    }
    ans = INF;
    for (int i = 0; i < n; i ++) {
        for (int j = 0; j < m[i]; j ++) {
            init();
            f[i][j][1 << i] = 0;
            calc(i, j);
        }
    }
    ans /= x;
    cout << fixed << setprecision(8) << ans + add << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t-- > 0) {
        solve();
    }
    return 0;
}
// 64 位输出请用 printf("%lld")