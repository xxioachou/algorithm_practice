#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int m;
    string s;
    cin >> s >> m;    
    int n = s.size();
    vector<int> idx;
    ll t = 0;
    for (int i = 0; i < n; i ++) {
        t = (t * 10 + s[i] - '0') % m;
    }
    if (t == 0) {
        cout << 0 << '\n';
        return;
    }

    auto print = [&]() {
        cout << "[";
        for (int i = 0; i < sz(idx); i ++) {
            cout << idx[i] << ","[i == sz(idx) - 1];
        }
        cout << "]\n";
    };
    int ans = inf;
    
    auto dfs = [&](auto dfs, int u) -> void {
        if (u == n - 1) {
            ll sum = 0;
            idx.push_back(n - 1);
            for (int i = 0, j = 0; j < sz(idx); j ++) {
                ll res = 0;
                while (i < n && i <= idx[j]) {
                    res = (res * 10 + s[i] - '0') % m;
                    i ++;
                }
                (sum += res) %= m;
            }
            idx.pop_back();

            if (sz(idx) && sum == 0) {
                // print();
                ans = min(ans, sz(idx));
            }
            return;
        }

        dfs(dfs, u + 1);

        idx.push_back(u);
        dfs(dfs, u + 1);
        idx.pop_back();
    };
    dfs(dfs, 0);

    if (ans >= inf) {
        ans = -1;
    }
    cout << ans << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}
