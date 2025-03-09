#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int finalPositionOfSnake(int n, vector<string>& commands) {
        int x = 0, y = 0;
        for (auto &str : commands) {
            if (str[0] == 'L') {
                y --;
            } else if (str[0] == 'R') {
                y ++;
            } else if (str[0] == 'U') {
                x --;
            } else {
                x ++;
            }
        }
        return x * n + y;
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