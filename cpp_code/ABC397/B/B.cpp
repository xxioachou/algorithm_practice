#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    string s;
    cin >> s;
    int turn = 0, ans = 0;
    for (int i = 0; i < s.size(); ) {
        if (!turn) {
            if (s[i] != 'i') {
                ans ++;
            } else {
                i ++;
            }
        } else {
            if (s[i] != 'o') {
                ans ++;
            } else {
                i ++;
            }
        }
        turn ^= 1;
        // cout << ":" << ans << '\n';
    }
    if ((s.size() + ans) % 2 == 1) {
        ans ++;
    }
    cout << ans << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}