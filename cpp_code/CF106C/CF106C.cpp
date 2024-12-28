#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, m, a0, b0, c0, d0;
    cin >> n >> m >> c0 >> d0;
    vector<tuple<int, int, int, int>> v;
    v.emplace_back(0, 0, c0, d0);
    for (int i = 0; i < m; i++) {
        cin >> a0 >> b0 >> c0 >> d0;
        v.emplace_back(a0, b0, c0, d0);
    }
    m = sz(v);
    vector<vector<int>> f(m + 1, vector<int>(n + 1));
    for (int i = 1; i <= m; i++) {
        auto &[a, b, c, d] = v[i - 1];
        for (int j = 0; j <= n; j++) {
            for (int k = 0; k * c <= j and k * b <= a; k++) {
                f[i][j] = max(f[i][j], f[i - 1][j - k * c] + k * d);
            }
        }
    }
    int ans = 0;
    for (int i = 0; i <= n; i++) {
        ans = max(ans, f[m][i]);
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