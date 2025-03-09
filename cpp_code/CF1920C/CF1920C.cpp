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
    vector<int> a(n);
    for (auto &x : a) {
        cin >> x;
    }
    int ans = 0;
    for (int k = 1; k <= n; k++) {
        if (n % k != 0)
            continue;

        int ok = 1, m = 0;
        for (int i = 0; i < k; i++) {
            for (int j = i + k; j < n; j += k) {
                m = __gcd(m, abs(a[j] - a[j - k]));
            }
        }
        if (m < 2) {
            if (m == 0)
                ans += 1;
            continue;
        }
        for (int i = 0; i < k and ok; i++) {
            int x = a[i] % m;
            for (int j = i; j < n; j += k) {
                if (a[j] % m != x) {
                    ok = 0;
                    break;
                }
            }
        }
        // if (ok)
        //     cout << ":" << k << ' ' << m << '\n';
        ans += ok;
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