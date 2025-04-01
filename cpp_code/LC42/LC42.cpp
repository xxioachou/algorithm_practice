#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int trap(vector<int>& height) {
        vector<pair<int, int>> p;
        vector<int> vis(height.size());
        long long tot = 0;
        for (int i = 0; i < height.size(); i ++) {
            tot -= height[i];
            p.emplace_back(height[i], i);
        }

        sort(p.begin(), p.end());
        long long last = 0, l = 0, r = height.size() - 1;
        for (int i = 0; i < p.size(); ) {
            int j = i + 1;
            while (j < p.size() && p[j].first == p[i].first) {
                j ++;
            }
            while (l < height.size() && vis[l]) {
                l ++;
            }
            while (r < height.size() && vis[r]) {
                r --;
            }

            tot += (p[i].first - last) * (r - l + 1);
            for (int k = i; k < j; k ++) {
                vis[p[k].second] = 1;
            }
            last = p[i].first;
            i = j;
        }
        return tot;
    }
};
