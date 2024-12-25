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
    vector<ll> a(n), b(n), c(n);
    for (auto &x : a) {
        cin >> x;
    }
    for (auto &x : b) {
        cin >> x;
    }
    
    ll ans = 0;
    iota(c.begin(), c.end(), 0);
    sort(c.begin(), c.end(), [&](const int &x, const int &y) {
        return a[x] < a[y];
    });
    vector<ll> bitmask;
    for (int i = 0; i < n; ) {
        int j = i + 1;
        while (j < n and a[c[j]] == a[c[i]]) {
            j ++;
        }
        if (j - i > 1) {
            bitmask.push_back(a[c[i]]);
        }
        i = j;
    }
    vector<int> vis(n);
    for (auto &bit : bitmask) {
        for (int i = 0; i < n; i++) {
            if ((a[i] & bit) == a[i]) {
                vis[i] = 1;
            }
        }
    }
    for (int i = 0; i < n; i++) {
        if (vis[i]) {
            ans += b[i];
        }
    }
    cout << ans << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    // freopen("6.in", "r", stdin);
    // freopen("CF1210B.out", "w", stdout);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}