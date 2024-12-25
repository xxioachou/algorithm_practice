#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<int> closestRoom(vector<vector<int>>& rooms, vector<vector<int>>& queries) {
        // 按照 size 从大到小排序
        sort(rooms.begin(), rooms.end(), [&](const vector<int> &a, const vector<int> &b) {
            return a[1] > b[1];
        });
        int n = rooms.size();
        int k = queries.size();
        const int inf = 0x3f3f3f3f;
        vector<int> idx(k), ans(k, -1);
        iota(idx.begin(), idx.end(), 0);
        // 按照 size 从大到小排序
        sort(idx.begin(), idx.end(), [&](const int &x, const int &y) {
            return queries[x][1] > queries[y][1];
        });
        set<int> st;

        for (int i = 0, j = 0; i < k; i++) {
            int p = queries[idx[i]][0];
            int sz = queries[idx[i]][1];
            while (j < n && rooms[j][1] >= sz) {
                st.emplace(rooms[j][0]);
                j ++;
            }
            if (st.empty()) {
                continue;
            }
            auto it = st.lower_bound(p);
            int r = (it == st.end() ? inf : *it);
            int l = (it == st.begin() ? -inf : *(--it));
            if (r - p >= p - l) {
                ans[idx[i]] = l;
            } else {
                ans[idx[i]] = r;
            }
        }
        return ans;
    }
};

void solve() {
    Solution s;
    
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}