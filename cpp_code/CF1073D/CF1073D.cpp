#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    ll n, T;
    cin >> n >> T;
    vector<ll> a(n), vis(n);
    vector<pair<ll, ll>> b;
    for (int i = 0; i < n; i++) {
        cin >> a[i];
        b.emplace_back(a[i], i);
    }
    sort(b.begin(), b.end());
    ll ans = 0, i = n - 1;
    ll sum = accumulate(a.begin(), a.end(), 0ll);
    while (true) {
        while (i >= 0 and T < b[i].x) {
            vis[b[i].y] = 1;
            sum -= b[i].x;
            i --;
        }
        if (i < 0)
            break;
        ll tmp = (T - b[i].x) / sum;
        ans += tmp * (i + 1);
        T -= tmp * sum;
        for (int j = 0; j < n; j++) {
            if (vis[j] or T < a[j])
                continue;
            T -= a[j];
            ans ++;
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