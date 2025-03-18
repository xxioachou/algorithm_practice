#include <algorithm>
#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int N = 2010;
const int mod = 1000000007;

ll qpow(ll a, ll b, ll mod) {
    ll res = 1 % mod;
    while (b) {
        if (b & 1) {
            res = res * a % mod;
        }
        b >>= 1;
        a = a * a % mod;
    }
    return res;
}

int p[N], cnt[N];

int find(int x) {
    return p[x] == x ? p[x] : p[x] = find(p[x]);
}

void merge(int x, int y) {
    // cout << x + 1 << ' ' << y + 1 << '\n';
    x = find(x);
    y = find(y);
    if (x != y) {
        p[x] = y;
        cnt[y] += cnt[x];
    }
}

void solve() {
    int n, m, k;
    cin >> n >> m >> k;
    iota(p, p + n, 0);
    fill(cnt, cnt + n, 1);
    for (int i = 0; i + k - 1 < n; i ++) {
        for (int j = i; j <= i + k - 1; j ++) {
            merge(j, i + i + k - 1 - j);
        }
    }
    int block = 0;
    for (int i = 0; i < n; i ++) {
        if (find(i) == i) {
            block ++;
        }
    }
    cout << qpow(m, block, mod) << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}