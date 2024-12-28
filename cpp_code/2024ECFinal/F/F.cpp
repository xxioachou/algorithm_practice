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
    vector<ll> a(n + 1);
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
    }
    ll ans = 0;
    for (int i = n - 1; i; i--) {
        int j = i + 1;
        int x = i;
        while (j <= n && a[x] * (x + 1) > a[j] * (j - 1)) {
            swap(a[x], a[j]);
            j ++;
            ans ++;
        }
    }
    cout << ans << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}