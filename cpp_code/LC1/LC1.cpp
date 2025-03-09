#include <bits/stdc++.h>
using namespace std;

class Solution {
    public:
        vector<int> twoSum(vector<int>& nums, int target) {
            std::unordered_map<int, int> mp;
            for (int i = 0; i < mp.size(); i ++) {
                if (mp.count(target - nums[i])) {
                    return std::vector<int>{i, mp[target - nums[i]]};
                }
                mp[nums[i]] = i;
            }
            return std::vector<int>{};
        }
};

void solve() {
    Solution s;
    string x;
    std::vector<int> v;
    while (std::cin >> x) {
        std::cout << x << ' ';
    }
    std::cout << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
//  cin >> t;
    while (t--) solve();
    return 0;
}