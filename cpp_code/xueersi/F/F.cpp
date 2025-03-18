#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, m, H;
    cin >> n >> m >> H;
    vector<map<int, int>> f(n + 1);
    f[0][0] = 0;
    for (int i = 1, c, d; i <= n; i ++) {
        cin >> c >> d;
        for (auto [j, v] : f[i - 1]) {
            for (int k = 0; j + k * c <= m; k ++) {
                f[i][j + k * c] = max(f[i][j + k * c], v + k * d);
            }
        }
    }
    for (auto [j, v] : f[n]) {
        if (v >= H) {
            cout << m - j << '\n';
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
