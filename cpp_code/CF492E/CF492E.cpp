#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int N = 1000010;

int mp[N];

void solve() {
    int n, m, dx, dy;
    cin >> n >> m >> dx >> dy;
    vector<int> x(m), y(m);
    auto get = [&](int x, int y) {
        int tmp =  1ll * x * dy % n - 1ll * y * dx % n;
        tmp = (tmp % n + n) % n;
        return tmp;
    };
    for (int i = 0; i < m; i++) {
        cin >> x[i] >> y[i];
        mp[get(x[i], y[i])] ++;
    }   
    int maxv = 0, a = -1, b = -1, t = -1;
    for (int i = 0; i < n; i++) {
        if (maxv < mp[i]) {
            maxv = mp[i];
            t = i;
        }
    }
    for (int i = 0; i < m; i++) {
        if (get(x[i], y[i]) == t) {
            cout << x[i] << ' ' << y[i] << '\n';
            return;
        }
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