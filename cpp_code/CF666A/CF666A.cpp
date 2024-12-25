#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    string s;
    cin >> s;
    int n = sz(s);
    vector<vector<int>> f(n + 5, vector<int>(2));
    vector<string> ans;
    f[n][0] = f[n][1] = 1;
    for (int i = n - 2; i >= 5; i--) {
        f[i][0] |= f[i + 2][1];
        if (f[i + 2][0] and s.substr(i, 2) != s.substr(i + 2, 2)) {
            f[i][0] = 1;
        }

        if (i + 3 <= n) {
            f[i][1] |= f[i + 3][0];
            if (f[i + 3][1] and s.substr(i, 3) != s.substr(i + 3, 3)) {
                f[i][1] = 1;
            }
        }

        if (f[i][0]) {
            ans.emplace_back(s.substr(i, 2));
        }
        if (f[i][1]) {
            ans.emplace_back(s.substr(i, 3));
        }
    }
    // for (int i = 0; i <= n; i++) {
    //     cout << f[i] << " \n"[i == n];
    // }
    sort(ans.begin(), ans.end());
    ans.erase(unique(ans.begin(), ans.end()), ans.end());
    cout << sz(ans) << '\n';
    for (auto &str : ans) {
        cout << str << '\n';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}