#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    const int inf = 0x3f3f3f3f;
    vector<int> minReverseOperations(int n, int p, vector<int>& banned, int k) {
        vector<int> vis(n), ans(n, inf);
        for (auto &x : banned) {
            vis[x] = 1;
        }
        
        vector<set<int>> st(2);
        for (int i = 0; i < n; i ++) {
            if (!vis[i]) {
                st[i % 2].emplace(i);
                // printf("st[%d]: %d\n", i % 2, i);
            }
        }
        

        queue<int> q;
        q.push(p);
        ans[p] = 0;
        st[0].erase(p);
        while (!q.empty()) {
            int u = q.front();
            q.pop();

            int l = max(0, u + 1 - k);
            int r = min(u, n - k);
            // printf("u: %d, l: %d, r: %d\n", u, l, r);
            int i = l, t = (k - 1 - u) % 2 != 0;
            // printf("i: %d, t: %d\n", i, t);
            while (i <= r) {
                int x = i * 2 + (k - 1 - u);
                auto it = st[t].lower_bound(x);
                if (it == st[t].end()) {
                    // printf("break, i: %d\n", i);
                    break;
                }

                int v = *it;
                int j = (v + u + 1 - k) / 2;
                if (j > r) {
                    break;
                }
                // printf("u: %d, v: %d\n", u, v);
                if (ans[v] > ans[u] + 1) {
                    ans[v] = ans[u] + 1;
                    q.push(v);
                    st[t].erase(it);
                }
                i = j + 1;
            }
        }
        for (auto &x : ans) {
            if (x >= inf) {
                x = -1;
            }
        }
        return ans;
    }
};
    
    
