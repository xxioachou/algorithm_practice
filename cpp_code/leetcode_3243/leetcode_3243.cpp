#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> shortestDistanceAfterQueries(int n, vector<vector<int>>& queries) {
        vector<int> dis(n), ans;
        vector<vector<int>> e(n);
        iota(dis.begin(), dis.end(), 0);
        for (int i = 0; i + 1 < n; i++) {
            e[i].push_back(i + 1);
        }
        for (auto &vec : queries) {
            int u = vec[0], v = vec[1];
            e[u].push_back(v);
            for (int i = 0; i < n; i++) {
                for (int x : e[i]) {
                    dis[x] = min(dis[x], dis[i] + 1);
                }
            }
            ans.emplace_back(dis[n - 1]);
        }
        return ans;
    }
};

void solve() {
    Solution s;
    int n, m;
    cin >> n >> m;
    vector<vector<int>> v;
    for (int i = 0, x, y; i < m; i++) {
        cin >> x >> y;
        v.emplace_back(vector<int>{x, y});
    }
    auto ans = s.shortestDistanceAfterQueries(n, v);
    for (int x : ans) {
        cout << x << ' ';
    }
    cout << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}