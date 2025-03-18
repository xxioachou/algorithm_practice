#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int u, v;
    cin >> u >> v;
    if (u > v) {
        cout << "NO\n";
        return;
    }

    int cnt1 = 0, cnt2 = 0;
    bool ok = true;
    for (int i = 0; i < 30; i ++) {
        int a = u >> i & 1;
        int b = v >> i & 1;
        if (a == b) {
            continue;
        }
        if (a == 1) {
            cnt1 ++;
        } else {
            cnt2 ++;
        }
        if (cnt1 < cnt2) {
            ok = false;
            // cout << "b:" << i << ' ' << cnt1 << ' ' << cnt2 << '\n';
            break;
        }
    }
    cout << (ok ? "YES\n" : "NO\n");
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}