#include <algorithm>
#include <bits/stdc++.h>
#define x first
#define y second
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int N;
    cin >> N;
    vector<vector<int>> e(N);
    vector<int> son(N), fa(N, -1);
    for (int i = 1, x, y; i < N; i ++) {
        cin >> x >> y;
        x --, y --;
        e[x].push_back(y);
        e[y].push_back(x);
    }
    int ans = inf;
    auto dfs1 = [&](auto dfs1, int u) -> void {
        // grandson_size
        for (int v : e[u]) {
            if (v != fa[u]) {
                fa[v] = u;
                dfs1(dfs1, v);
                son[u] ++;
            }
        }
    };
    auto dfs2 = [&](auto dfs2, int u) -> void {
        // grandson_size
        vector<int> gs, suf;
        for (int v : e[u]) {
            if (v != fa[u]) {
                dfs2(dfs2, v);
                gs.push_back(son[v]);
            }
        }
        if (fa[u] != -1) {
            gs.push_back(son[fa[u]] - 1 + (fa[fa[u]] != -1 ? 1 : 0));
        }

        int n = gs.size();
        sort(gs.begin(), gs.end());
        // cout << u + 1 << "\ngs: ";
        // for (int x : gs) {
        //     cout << x << ' ';
        // }
        // cout << '\n';
        suf.resize(n + 1);
        for (int i = n - 1; i >= 0; i --) {
            suf[i] = suf[i + 1] + gs[i];
        }

        int pre = 0;
        int tot = n + 1 + suf[0];
        for (int i = 0, j = 1; i < n; ) {
            while (j < n && gs[j] == gs[i]) {
                j ++;
            }

            int cost = i + pre + suf[j] - (n - j) * gs[i];
            cost += N - tot;
            ans = min(ans, cost);
            pre += (j - i) * gs[i];
            i = j;
        }

    };
    dfs1(dfs1, 0);
    dfs2(dfs2, 0);
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
