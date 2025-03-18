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
    map<int, set<int>> v;
    int ans = 0;
    while (n -- > 0) {
        int op, price, t;
        cin >> op >> price >> t;

        if (!op) {
            ans += price;
            v[t].insert(price);
        } else {
            bool ok = false;
            for (int i = t - 45; i <= t; i ++) {
                if (v.find(i) == v.end()) {
                    continue;
                }
                auto it = v[i].lower_bound(price);
                if (it != v[i].end()) {
                    ok = true;
                    v[i].erase(it);
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