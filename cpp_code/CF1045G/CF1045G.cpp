#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int MAXN = 100010;

struct node {
    map<int, int> mp;
} tr[MAXN * 4];

void update(int u, int l, int r, int x, int v) {
    if (l == r) {
        tr[u].mp[v] += 1;
        return;
    }
    int mid = (l + r) / 2;
    if (l <= mid) 
        update(u << 1, l, mid, x, v);
    else
        update(u << 1 | 1, mid + 1, r, x, v);
    pushup(u);
}

int query(int u, int l, int r, int L, int R, int q) {

}

void solve() {
    int N, K;
    cin >> N >> K;
    vector<tuple<int, int, int>> a;
    vector<int> v; 
    for (int i = 0, x, r, q; i < N; i++) {
        cin >> x >> r >> q;
        a.emplace_back(x, r, q);
        v.emplace_back(x + r);
    }
    sort(v.begin(), v.end());
    sort(a.begin(), a.end());
    v.erase(unique(v.begin(), v.end()), v.end());
    auto get = [&](int x) {
        return lower_bound(v.begin(), v.end(), x) - v.begin() + 1;
    };
    int m = v.size();
    ll ans = 0;
    // [1, m]
    for (auto [x, r, q] : a) {
        auto it = lower_bound(v.begin(), v.end(), x - r);
        if (it == v.end())
            continue;
        int idx = it - v.begin() + 1;
        ans += query(1, 1, m, idx, m, q);
        update(1, 1, m, get(x + r), q);
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