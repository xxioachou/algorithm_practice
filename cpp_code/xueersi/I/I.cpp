#include <bits/stdc++.h>
#include <ios>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

struct bitTree {
    int n;
    vector<int> a;
    bitTree() {}
    bitTree(int sz) {
        n = sz;
        a.resize(sz + 1);
    }
    int lowbit(int x)   { return x &-x; }
    void add(int x, int v) {
        while (x <= n) {
            a[x] += v;
            x += lowbit(x);
        }
    }
    int sum(int x) {
        int res = 0;
        while (x > 0) {
            res += a[x];
            x -= lowbit(x);
        }
        return res;
    }
};

void solve() {
    int k, n, m;
    cin >> k >> n >> m;
    vector<int> vis(n), c(n), a(m), p;
    for (int i = 0; i < m; i ++) {
        cin >> a[i];
        p.push_back(i);
    }
    for (int i = 0, cnt = 0, cur = 0; cnt < n - 1; cur = (cur + 1) % n) {   
        while (vis[cur]) {
            cur = (cur + 1) % n;
        }

        cout << cur + 1 << ' ' << p[i] + 1 << ' ' << a[p[i]] << '\n';
        if (a[p[i]] != 0) {
            if (a[p[i]] == 1) {
                c[cur] ++;
                i ++;
            } else {
                if (c[cur] == 0) {
                    vis[cur] = 1;
                    i ++;
                    cnt ++;
                } else {
                    c[cur] --;
                    
                    int t = p[i];
                    for (int j = i + 1; j <= i + k && j < m; j ++) {
                        p[j - 1] = p[j];
                    }
                    p[i + k] = t;
                }
            }
        } else {
            i ++;
        }
    }
    for (int i = 0; i < n; i ++) {
        if (!vis[i]) {
            cout << i + 1 << '\n';
            return;
        }
    }
    // 2 1 0 1 1 0 0 0 0 1 0 1
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}
