#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n, pre = 0;
    string s;
    cin >> n >> s;
    array<int, 4> cnt;
    fill(cnt.begin(), cnt.end(), 0);

    for (auto &ch : s) {
        pre ^= (ch == '0' ? 2 : 1);
        cnt[pre] ++;
    }

    pre = 0;
    int res = 0;
    for (int i = 0; i < n; i ++) {
        pre ^= (s[i] == '0' ? 2 : 1);
        cnt[pre] --;
        if (cnt[pre] > 0) {
            res ++;
        }
    }
    cout << fixed << setprecision(8) << 1.0 * res / (1.0 * n) << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}
