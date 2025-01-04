#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, q;
    cin >> n >> q;
    vector s(n + 1, vector(n + 1, 0ll));
    vector p(n + 1, vector(n + 1, 0ll));
    vector t(n + 1, vector(n + 1, 0ll));
    auto sum = [&](vector<vector<ll>> &v, int x1, int y1, int x2, int y2) -> ll {
        return v[x2][y2] - v[x1 - 1][y2] - v[x2][y1 - 1] + v[x1 - 1][y1 - 1];
    };
    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= n; j++) {
            ll a;
            cin >> a;
            s[i][j] = s[i][j - 1] + s[i - 1][j] - s[i - 1][j - 1] + a;
            p[i][j] = p[i][j - 1] + p[i - 1][j] - p[i - 1][j - 1] + a * i;
            t[i][j] = t[i][j - 1] + t[i - 1][j] - t[i - 1][j - 1] + a * j;
        }
    }
    for (int i = 0, x1, y1, x2, y2; i < q; i++) {
        cin >> x1 >> y1 >> x2 >> y2;
        auto sum1 = sum(s, x1, y1, x2, y2);
        auto ans = sum(t, x1, y1, x2, y2) - sum1 * (y1 - 1);
        auto tmp = sum(p, x1, y1, x2, y2) - sum1 * x1;
        ans += tmp * (y2 - y1 + 1);
        cout << ans << ' ';
    }
    cout << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}