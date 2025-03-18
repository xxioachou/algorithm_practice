#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int mod = 998244353;

int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

void solve() {
    int h, w, sx, sy;
    cin >> h >> w >> sx >> sy;
    sx --, sy --;
    vector<string> mp(h);
    vector<vector<int>> f(h, vector<int>(w));
    vector<vector<int>> dis(h, vector<int>(w, inf));
    for (auto &s : mp) {
        cin >> s;
    }
    {
        queue<pair<int, int>> q;
        q.emplace(sx, sy);
        f[sx][sy] = 1;
        dis[sx][sy] = 0;
        while (sz(q) > 0) {
            auto [x, y] = q.front();
            q.pop();

            for (int i = 0; i < 4; i ++) {
                int a = x + dx[i];
                int b = y + dy[i];
                if (a < 0 || a >= h || b < 0 || b >= w || mp[a][b] == '#' || dis[a][b] < dis[x][y] + 1) {
                    continue;
                }
                if (dis[a][b] > dis[x][y] + 1) {
                    dis[a][b] = dis[x][y] + 1;
                    f[a][b] = f[x][y];
                    q.emplace(a, b);
                } else if (dis[a][b] == dis[x][y] + 1) {
                    f[a][b] += f[x][y];
                    f[a][b] %= mod;
                    // q.emplace(a, b);
                }
            }
        }
    }
    int q;
    cin >> q;
    for (int i = 0, ex, ey; i < q; i ++) {
        cin >> ex >> ey;
        ex --, ey --;
        cout << f[ex][ey] << '\n';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}