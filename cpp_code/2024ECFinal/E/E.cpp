#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    ll n, d, N = 0;
    cin >> n >> d;
    vector<ll> a(n + 1);
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        N += a[i];
    }
    sort(a.begin() + 1, a.end());
    ll ans = 0;
    if (d == 1) {
        ans = 1ll * N * (N - 1) / 2;
    } else if (d == 2) {
        for (int i = n; i; i--) {
            N -= a[i];
            ans += a[i] * N;
        }
    } else {
        for (int i = n; i; i--) {
            N -= a[i];
            ans += N;
        }
    }
    cout << ans << '\n';
    // int maxv = *max_element(a.begin() + 1, a.end());
    // int minv = *min_element(a.begin() + 1, a.end());
    // if (d == 1 || maxv <= 1) {
    //     cout << 1ll * N * (N - 1) / 2 << '\n';
    // } else {
    //     if (maxv > N - maxv) {
    //         cout << 0 << '\n';
    //     } else {
    //         int t = (N % d == 0 ? 1 : 0);
    //         cout << N - 1 + t << '\n';
    //     }
    // }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}