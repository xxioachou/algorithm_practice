#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, m, k;
    cin >> n >> m >> k;
    set<pair<int, int>> st;
    vector<int> x(n + 1), B(m + 1);
    vector<int> y(n + 1, inf);
    for (int i = 1, a, b; i <= m; i ++) {
        if (i - k - 1 > 0) {
            st.erase(st.find(pair<int, int>{B[i - k - 1], i - k - 1}));
        }

        cin >> a >> b;
        B[i] = b;

        int v = B[i];
        if (sz(st) >= 2) { 
            auto it1 = prev(st.end());
            auto it2 = prev(it1);
            v = (it1->first) / 10 + (it2->first) / 5 + b; 
        }

        x[a] ++;
        y[a] = min(y[a], v);
        st.emplace(b, i);
        cout << a << ' ' << i << ' ' << v << '\n';

    }
    for (int i = 1; i <= n; i ++) {
        cout << 1ll * x[i] * y[i] * y[i] << '\n';
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
