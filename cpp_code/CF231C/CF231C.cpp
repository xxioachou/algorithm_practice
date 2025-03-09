#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, k;
    cin >> n >> k;
    vector<int> a(n);
    for (auto &x : a) {
        cin >> x;
    }
    sort(a.begin(), a.end());
    ll sum = 0;
    int ans1 = 0, ans2 = 0;
    for (int i = 0, j = 0; i < n; i++) {
        sum += a[i];
        while (j < i and 1ll * a[i] * (i - j + 1) - sum > k) {
            sum -= a[j ++];
        }
        
        if (1ll * a[i] * (i - j + 1) - sum <= k) {
            if (ans1 < i - j + 1) {
                ans1 = i - j + 1;
                ans2 = a[i];
            }
        }
    }
    cout << ans1 << ' ' << ans2 << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}