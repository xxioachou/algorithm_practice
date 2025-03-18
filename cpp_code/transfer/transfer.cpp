#include <bits/stdc++.h>
using namespace std;

void solve() {
    int n;
    cin >> n;
    deque<pair<int, int>> q;
    int ans = 0;
    while (n -- > 0) {
        int op, price, t;
        cin >> op >> price >> t;

        if (!op) {
            ans += price;
            q.emplace_back(t, price);
        } else {
            while (!q.empty() && t - q.front().first > 45) {
                q.pop_front();
            }
            bool ok = false;
            for (auto it = q.begin(); it != q.end(); it ++) {
                if (it->second >= price) {
                    ok = true;
                    break;
                }
            }
            if (!ok) {
                ans += price;
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