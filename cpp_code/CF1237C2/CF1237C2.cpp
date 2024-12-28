#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n;
    cin >> n;
    map<int, map<int, vector<pair<int, int>>>> mp;
    vector<int> ans;
    for (int i = 0, x, y, z; i < n; i++) {
        cin >> x >> y >> z;
        mp[x][y].emplace_back(z, i);
    }
    map<int, vector<pair<int, int>>> p;
    for (auto &[x, mmp] : mp) {
        for (auto &[y, v] : mmp) {
            sort(v.begin(), v.end());
            while (sz(v) > 1) {
                ans.emplace_back(v.back().y);
                v.pop_back();
                ans.emplace_back(v.back().y);
                v.pop_back();            
            }
            if (sz(v)) {
                p[x].emplace_back(y, v.back().y);
            }
        }
    }
    vector<int> tmp;
    for (auto &[x, v] : p) {
        sort(v.begin(), v.end());
        while (sz(v) > 1) {
            ans.emplace_back(v.back().y);
            v.pop_back();
            ans.emplace_back(v.back().y);
            v.pop_back();
        }
        if (sz(v)) {
            tmp.push_back(v.back().y);
            if (sz(tmp) > 1) {
                ans.emplace_back(tmp.back());
                tmp.pop_back();
                ans.emplace_back(tmp.back());
                tmp.pop_back();
            }
        }
    }
    for (int i = 0; i < n; i += 2) {
        cout << ans[i] + 1 << ' ' << ans[i + 1] + 1 << '\n';
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