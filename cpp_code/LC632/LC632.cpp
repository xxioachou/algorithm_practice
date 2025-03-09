#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    using node = tuple<int, int, int>;
    vector<int> smallestRange(vector<vector<int>>& nums) {
        priority_queue<node, vector<node>, greater<node>> q;
        vector<int> a;
        for (auto &v : nums) {
            for (auto &x : v) {
                a.push_back(x);
            }
        }
        sort(a.begin(), a.end());
        int R = -100010;
        for (int i = 0; i < nums.size(); i++) {
            q.emplace(nums[i][0], i, 0);
            R = max(R, nums[i][0]);
        }
        int minLen = 200010, l = -1, r = -1;
        for (int x : a) {
            while (!q.empty()) {
                auto [e, i, j] = q.top();
                if (e >= x) {
                    break;
                }
                q.pop();
                if (j + 1 < nums[i].size()) {
                    q.emplace(nums[i][j + 1], i, j + 1);
                    R = max(R, nums[i][j + 1]);
                }
            }
            if (q.size() >= nums.size()) {
                if (R - x < minLen) {
                    minLen = R - x;
                    l = x, r = R;
                }
            }
        }
        return vector<int>{l, r};
    }
};

void solve() {
    Solution s;
    
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
//  cin >> t;
    while (t--) solve();
    return 0;
}