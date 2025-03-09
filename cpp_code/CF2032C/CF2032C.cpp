#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n;
    cin >> n;
    vector<int> a(n);
    for (auto &x : a) {
        cin >> x;
    }
    int ans = inf;
    sort(a.begin(), a.end());
    for (int i = 0, j = 1; i < n; i++) {
        while (j < n and a[i] + a[i + 1] > a[j]) {
            j ++;
        }
        int tmp = i;
        if (j < n and a[i] + a[i + 1] <= a[j]) {
            tmp += n - j;
        }
        ans = min(ans, tmp);
    }
    cout << ans << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}