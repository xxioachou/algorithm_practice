#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> shortestDistanceAfterQueries(int n, vector<vector<int>>& queries) {
        list<int> l;
        vector<int> ans;
        vector<list<int>::iterator> mp(n);        
        for (int i = 0; i < n; i++) {
            l.push_back(i);
            mp[i] = prev(l.end());
        }
        for (auto &vec : queries) {
            int u = vec[0], v = vec[1];
            if (mp[u] != l.end() and mp[v] != l.end()) {
                for (auto it = next(mp[u]); it != mp[v]; ) {
                    mp[*it] = l.end();
                    auto tmp = it;
                    it++;
                    l.erase(tmp);
                }
            }
            ans.emplace_back(l.size() - 1);
        }
        return ans;
    }
};

void solve() {
    Solution s;
    int n, m;
    cin >> n >> m;
    vector<vector<int>> q;
    for (int i = 0, x, y; i < m; i++) {
        cin >> x >> y;
        q.emplace_back(vector<int>{x, y});
    }
    auto ans = s.shortestDistanceAfterQueries(n, q);
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