#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    ll N;
    cin >> N;
    auto get = [&](ll d) -> ll {
        ll l = 1, r = 1000'000'000;
        while (l < r) {
            ll mid = (l + r) / 2;
            if (mid * mid + d * mid >= (N - d * d * d + d * 3ll - 1) / (d * 3ll)) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        return r;
    };
    for (ll d = 1; d * d * d <= N; d ++) {
        ll y = get(d);
        if (3ll * d * y * y + 3ll * d * d * y == N - d * d * d) {
            cout << y + d << ' ' << y << '\n';
            return;
        }
    }
    cout << -1 << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}