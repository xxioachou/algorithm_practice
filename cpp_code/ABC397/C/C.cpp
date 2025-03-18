#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int N;
    cin >> N;
    vector<int> A(N), pre(N + 1), suf(N + 1);
    int p = 0, s = 0;
    for (auto &x : A) {
        cin >> x;
        if (++ suf[x] == 1) {
            s ++;
        }
    }
    if (-- suf[A[0]] == 0) {
        s --;
    }
    if (++ pre[A[0]] == 1) {
        p ++;
    }
    int ans = p + s;
    for (int i = 1; i + 1 < N; i ++) {
        if (-- suf[A[i]] == 0) {
            s --;
        }
        if (++ pre[A[i]] == 1) {
            p ++;
        }
        ans = max(ans, p + s);
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