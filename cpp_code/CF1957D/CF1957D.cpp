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
    vector<int> a(n + 1), s(n + 1);
    vector<int> cnt0(30), cnt1(30);
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        s[i] = s[i - 1] ^ a[i];
        for (int j = 0; j < 30; j++) {
            if (s[i] >> j & 1)
                cnt1[j] += 1;
            else
                cnt0[j] += 1;
        }
    }
    vector<int> pre0(30, 1), pre1(30);
    ll ans = 0;
    for (int y = 1; y <= n; y++) {
        for (int j = 29; j >= 0; j--) {
            if (a[y] >> j & 1) {
                ans += 1ll * pre0[j] * cnt0[j];
                ans += 1ll * pre1[j] * cnt1[j];
                break;
            }
        }

        for (int j = 0; j < 30; j++) {
            if (s[y] >> j & 1) {
                cnt1[j] -= 1;
                pre1[j] += 1;
            } else {
                cnt0[j] -= 1;
                pre0[j] += 1;
            }
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