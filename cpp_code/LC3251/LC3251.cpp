#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countOfPairs(vector<int>& nums) {
        int n = nums.size();
        int m = *max_element(nums.begin(), nums.end());
        const int mod = 1e9 + 7;
        vector<vector<int>> f(2, vector<int>(m + 1));
        vector<int> s(m + 2);
        for (int j = 0; j <= nums[0]; j++) {
            f[1][j] = 1;
        }
        for (int j = 0; j <= m; j++) {
            s[j + 1] = s[j];
            if (j <= nums[0]) {
                s[j + 1] = (s[j + 1] + f[1][j]) % mod;
            }
        }
        for (int i = 2; i <= n; i++) {
            for (int j = 0; j <= nums[i - 1]; j++) {
                f[i & 1][j] = 0;
            }
            for (int j = 0; j <= nums[i - 1]; j++) {
                int k = nums[i - 1] - j;
                int x = nums[i - 2];
                int l = max(x - j, k);
                int r = min(x, m);
                int sum = 0;
                if (l <= r)
                    sum = ((s[r + 1] - s[l]) % mod + mod) % mod;
                f[i & 1][k] = (f[i & 1][k] + sum) % mod;
            }
            for (int j = 0; j <= m; j++) {
                s[j + 1] = s[j];
                if (j <= nums[i - 1])
                    s[j + 1] = (s[j + 1] + f[i & 1][j]) % mod;
            }   
        }
        int ans = 0;
        for (int j = 0; j <= nums[n - 1]; j++) {
            ans = (ans + f[n & 1][j]) % mod;
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