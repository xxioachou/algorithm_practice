#include <bits/stdc++.h>
using namespace std;

void solve() {
    int n;
    cin >> n;
    int ans = 0;
    vector<int> op(n), price(n), t(n);
    for (int i = 0; i < n; i ++) {
        cin >> op[i] >> price[i] >> t[i];
    }

    // 存储所有可用的优惠
    vector<int> idx, vis(n);
    for (int i = 0; i < n; i ++) {
        if (!op[i]) {
            ans += price[i];
            idx.push_back(i);
        } else {
            bool ok = false;
            for (int j : idx) {
                if (price[j] >= price[i] && !vis[j] && t[i] - t[j] <= 45) {
                    vis[j] = 1;
                    ok = true;
                    break;
                }
            }
            if (!ok) {
                ans += price[i];
            }
        }
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