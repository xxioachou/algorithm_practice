#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, d, N = 0;
    cin >> n >> d;
    vector<int> a(n + 1);
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        N += a[i];
    }
    vector<int> c(N + 1);
    vector<vector<int>> st(n + 1);
    int idx = 0;
    for (int i = 1; i <= n; i++) {
        for (int j = 0; j < a[i]; j++) {
            idx ++;
            c[idx] = i;
            st[i].push_back(idx);
        }
    }
    vector<vector<int>> e(N + 1);
    auto check = [&] {
        vector<vector<int>> f(N + 1, vector<int>(N + 1, inf));
        for (int i = 1; i <= N; i++) {
            for (int j : e[i]) {
                f[i][j] = min(f[i][j], 1);
            }
        }
        for (int k = 1; k <= N; k++) {
            for (int i = 1; i <= N; i++) {
                for (int j = 1; j <= N; j++) {
                    f[i][j] = min(f[i][j], f[i][k] + f[k][j]);
                }
            }
        }
        for (int i = 1; i <= n; i++) {
            for (int x : st[i]) {
                for (int y : st[i]) {
                    if (y != x && f[x][y] < d)
                        return false;
                }
            }
        }
        return true;
    };
    int ans = 0;
    for (int i = 1; i <= N; i++) {
        for (int j = i + 1; j <= N; j++) {
            e[i].push_back(j);
            e[j].push_back(i);
            if (!check()) {
                e[i].pop_back();
                e[j].pop_back();
            } else {
                ans ++;
                cout << i << ' ' << j << ' ' << c[i] << ' ' << c[j] << '\n';
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