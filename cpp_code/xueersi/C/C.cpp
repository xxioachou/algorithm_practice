#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

// [0, m * 6] -> [1, m * 6 + 1]
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
    int n, m;
    cin >> n >> m;
    bitTree tr(m * 6 + 10);
    vector<int> sum(n + 1, 1);
    auto query = [&](int x) {
        int low = 1, high = m * 6 + 10;
        while (low < high) {
            int mid = (low + high) / 2;
            if (tr.sum(mid) >= x) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        return high;
    };
    auto R = [&]() -> int {
        return query(n);
    };
    auto Mid = [&]() -> int {
        return query((n + 1) / 2);
    };
    auto L = [&]() -> int {
        return query(1);
    };

    tr.add(1, n);
    for (int i = 0, id, x; i < m; i ++) {
        cin >> id >> x;
        tr.add(sum[id], -1);
        sum[id] += x;
        tr.add(sum[id], 1);

        cout << R() - 1 << ' ' << Mid() - 1 << ' ' << L() - 1 << '\n';
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