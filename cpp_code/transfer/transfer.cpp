#include <bits/stdc++.h>
using namespace std;

void solve() {
    int n;
    cin >> n;
    deque<int> q;
    vector<int> op(n), price(n), t(n), vis(n);
    for (int i = 0; i < n; i ++) {
        cin >> op[i] >> price[i] >> t[i];
    }
    int ans = 0;
    for (int i = 0; i < n; i ++) {
        if (op[i] == 0) {
            ans += price[i];
            q.push_back(i);
        } else {
            while (!q.empty() && t[i] - t[q.front()] > 45) {
                q.pop_front();
            }
            bool ok = false;
            for (auto it = q.begin(); it != q.end(); it ++) {
                if (!vis[*it] && price[*it] >= price[i]) {
                    vis[*it] = 1;
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