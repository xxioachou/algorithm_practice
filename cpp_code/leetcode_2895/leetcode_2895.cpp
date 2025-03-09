#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minProcessingTime(vector<int>& processorTime, vector<int>& tasks) {
        int n = processorTime.size();
        sort(processorTime.begin(), processorTime.end());
        sort(tasks.rbegin(), tasks.rend());
        int ans = 0;
        for (int i = 0; i < n * 4; i++) {
            ans = max(ans, tasks[i] + processorTime[i / 4]);
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