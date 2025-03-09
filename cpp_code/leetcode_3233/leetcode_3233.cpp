#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    bool is_pri(int x) {
        if (x < 2)
            return false;
        for (int i = 2; i <= x / i; i++) {
            if (x % i == 0)
                return false;
        }
        return true;
    }
    int nonSpecialCount(int l, int r) {
        int x = ceil(sqrtl(l));
        int y = floor(sqrtl(r));
        int ans = r - l + 1;
        for (; x <= y; x++) {
            if (is_pri(x)) {
                ans --;
            }
        }
        return ans;
    }
};

void solve() {
    Solution s;
    int l, r;
    cin >> l >> r;
    cout << s.nonSpecialCount(l, r) << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}