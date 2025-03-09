#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int networkDelayTime(vector<vector<int>>& times, int n, int k) {
        vector<vector<pair<int, int>>> e(n + 1);
        for (auto &v : times) {
            int x = v[0], y = v[1], t = v[2];
            e[x].emplace_back(y, t);
        }
        const int inf = 0x3f3f3f3f;
        vector<int> dis(n + 1, inf), vis(n + 1);
        dis[k] = 0;
        for (int i = 0; i < n; i++) {
            int t = -1;
            for (int j = 1; j <= n; j++) {
                if (vis[j])
                    continue;
                if (t == -1 or dis[t] > dis[j]) {
                    t = j;
                }
            }
            vis[t] = 1;
            for (auto [j, w] : e[t]) {
                dis[j] = min(dis[j], dis[t] + w);
            }
        }
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            ans = max(ans, dis[i]);
        }
        if (ans >= inf)
            ans = -1;
        return ans;
    }
};

void solve() {
    Solution s;
    int n, m, k;
    cin >> n >> m >> k;
    vector<vector<int>> t(m);
    for (auto &v : t) {
        int x, y, z;
        cin >> x >> y >> z;
        v = vector<int>{x, y, z};
    }
    cout << s.networkDelayTime(t, n, k) << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}