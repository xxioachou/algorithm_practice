#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int N = 400010;

struct node {
    int v;
} tr[N * 4];

void pushup(int u) {
    tr[u].v = max(tr[u << 1].v, tr[u << 1 | 1].v);
}

void build(int u, int l, int r) {
    tr[u] = {0};
    if (l == r) {
        return;
    }
    int mid = (l + r) / 2;
    build(u << 1, l, mid);
    build(u << 1 | 1, mid + 1, r);
}

void update(int u, int l, int r, int x, int v) {
    if (l == r) {
        tr[u].v = max(tr[u].v, v);
        return;
    }
    int mid = (l + r) / 2;
    if (x <= mid) {
        update(u << 1, l, mid, x, v);
    } else {
        update(u << 1 | 1, mid + 1, r, x, v);
    }
    pushup(u);
}

int query(int u, int l, int r, int L, int R) {
    if (L <= l && r <= R) {
        return tr[u].v;
    }
    int mid = (l + r) / 2;
    int res = 0;
    if (L <= mid) {
        res = max(res, query(u << 1, l, mid, L, R));
    }
    if (R > mid) {
        res = max(res, query(u << 1 | 1, mid + 1, r, L, R));
    }
    return res;
}

void solve() {
    int n;
    cin >> n;
    vector<int> l(n), r(n), point, ans(n);
    map<int, vector<int>> add, del;
    for (int i = 0; i < n; i++) {
        cin >> l[i] >> r[i];
        point.push_back(l[i]);
        point.push_back(r[i]);
        add[l[i]].push_back(i);
        del[r[i]].push_back(i);
    }
    
    sort(point.begin(), point.end());
    point.erase(unique(point.begin(), point.end()), point.end());
    auto get = [&](int x) {
        return lower_bound(point.begin(), point.end(), x) - point.begin() + 1;
    };
    int m = point.size();
    build(1, 1, m);
    for (int x : point) {
        for (int i : add[x]) {
            update(1, 1, m, get(r[i]), r[i] - l[i] + 1);
        }
        for (int i : add[x]) {
            ans[i] = query(1, 1, m, get(r[i]), m) - r[i] + l[i] - 1;
        }
        for (int i : del[x]) {
            update(1, 1, m, get(r[i]), 0);
        }
    }
    for (int x : ans) {
        cout << x << '\n';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}