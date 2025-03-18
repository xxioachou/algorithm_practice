#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int N = 100010;

int idx;
int root[N], dfn[N], out[N], p[N];
struct node {
    int l, r, cnt;
} tr[N * 4 + N * 17];

int insert(int p, int l, int r, int x) {
    int q = ++ idx;
    tr[q] = tr[p];
    if (l == r) {
        tr[q].cnt ++;
        return q;
    }
    int mid = (l + r) / 2;
    if (x <= mid) {
        tr[q].l = insert(tr[p].l, l, mid, x);
    } else {
        tr[q].r = insert(tr[p].r, mid + 1, r, x);
    }
    tr[q].cnt = tr[tr[q].l].cnt + tr[tr[q].r].cnt;
    return q;
}

int query(int p, int l, int r, int L, int R) {
    if (L <= l && r <= R) {
        return tr[p].cnt;
    }

    int mid = (l + r) / 2;
    int res = 0;
    if (L <= mid) {
        res += query(tr[p].l, l, mid, L, R);
    }
    if (R > mid) {
        res += query(tr[p].r, mid + 1, r, L, R);
    }
    return res;
}

void solve() {
    int n, q, timestmp = 0;
    cin >> n >> q;
    vector<vector<int>> e(n + 1);
    for (int i = 1, u, v; i < n; i ++) {
        cin >> u >> v;
        e[u].push_back(v);
        e[v].push_back(u);
    }
    for (int i = 1; i <= n; i ++) {
        cin >> p[i];
    }
    auto dfs = [&](auto dfs, int u, int fa) -> void {
        dfn[u] = ++ timestmp;
        for (int v : e[u]) {
            if (v != fa) {
                dfs(dfs, v, u);
            }
        }
        out[u] = timestmp;
    };
    dfs(dfs, 1, 0);

    idx = 0;
    for (int i = 1; i <= n; i ++) {
        root[i] = insert(root[i - 1], 1, n, dfn[p[i]]);
    }
    while (q -- > 0) {
        int l, r, x;
        cin >> l >> r >> x;
        int cnt = query(root[r], 1, n, dfn[x], out[x]) - 
            query(root[l - 1], 1, n, dfn[x], out[x]);
        cout << (cnt > 0 ? "YES\n": "NO\n");
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
