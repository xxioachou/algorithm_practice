#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, k;
    cin >> n >> k;
    vector<vector<int>> e(n + 1);
    for (int i = 1, x, y; i < n; i ++) {
        cin >> x >> y;
        e[x].push_back(y);
        e[y].push_back(x);
    }
    auto dfs = [&](auto dfs, int u, int fa, int x, int &cost) -> int {
        int cnt = 1;
        for (int v : e[u]) {
            if (v != fa) {
                int res = dfs(dfs, v, u, x, cost);
                cnt += res;
            }
        }
        if (cnt > x) {
            cost ++;
            return 0;
        }

        return cnt;
    };
    int cost = 0;
    int l = 0, r = n;
    while (l < r) {
        int mid = (l + r) / 2;
        
        cost = 0;
        dfs(dfs, 1, 0, mid, cost);
        if (cost <= k) {
            r = mid;
        } else {
            l = mid + 1;
        }
    }
    cout << r << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}
