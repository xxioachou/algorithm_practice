#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        int n = s1.size();
        int m = s2.size();
        if (m < n)
            return false;
        vector<int> cnt1(26), cnt2(26);
        for (auto &x : s1) {
            cnt1[x - 'a'] += 1;
        }
        queue<int> q;
        for (int i = 0; i < m; i++) {
            if (!q.empty() and q.front() < i - n + 1) {
                int x = s2[q.front()] - 'a';
                cnt2[x] --;
                q.pop();
            }
            q.push(i);
            cnt2[s2[i] - 'a'] ++;
            if (q.size() >= n and cnt1 == cnt2) {
                return true;
            }
        }
        return false;
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