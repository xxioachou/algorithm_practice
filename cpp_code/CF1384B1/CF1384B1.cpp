#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, k, l;
    cin >> n >> k >> l;
    int m = k * 2;
    vector<int> d(n + 5), p(m);
    for (int i = 1; i <= n; i++) {
        cin >> d[i];
    }
    for (int i = 0; i < m; i++) {
        if (i < k) {
            p[i] = i;
        } else {
            p[i] = k - (i - k);
        }
    }
    vector<vector<int>> f(n + 5, vector<int>(m));
    f[0][0] = 1;
    d[0] = d[n + 1] = -inf;
    for (int i = 0; i <= n; i++) {
        for (int j = 0; j < m; j++) {
            for (int t = 1; t <= m; t++) {
                int v = (j + t) % m;
                int u = (v - 1 + m) % m; 
                if (d[i] + p[u] > l) 
                    break;
                if (d[i + 1] + p[v] <= l) {
                    f[i + 1][v] |= f[i][j];
                }
            }
            // cout << i << ' ' << j << ' ' << f[i][j] << '\n';
        }
    }
    for (int x : f[n + 1]) {
        if (x) {
            cout << "Yes\n";
            // cout << x << '\n';
            return;
        }
    }
    cout << "No\n";
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}