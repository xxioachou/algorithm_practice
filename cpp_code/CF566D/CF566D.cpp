#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int N = 200010;

int ne[N], p[N];

int find(int x) {
    return p[x] == x ? p[x] : p[x] = find(p[x]);
}

void merge(int x, int y) {
    x = find(x);
    y = find(y);
    if (x != y) {
        p[x] = y;
    }
}

void solve() {
    int n, q;
    cin >> n >> q;
    iota(p + 1, p + 1 + n, 1);
    iota(ne + 1, ne + 1 + n, 2);
    for (int i = 1, t, x, y; i <= q; i++) {
        cin >> t >> x >> y;
        if (t == 1) {
            merge(x, y);
        } else if (t == 2) {
            for (int i = x, last = -1; i <= y; ) {
                if (last != -1) {
                    merge(last, i);
                }
                int t = ne[i];
                ne[i] = ne[y];
                last = i;
                i = t;
            }
        } else {
            cout << (find(x) == find(y) ? "YES\n" : "NO\n");
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