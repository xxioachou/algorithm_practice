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
    int ok = 1;
    map<char, int> mp;
    vector<int> f(3);
    for (char c : s) {
        mp[c]++;
        if (isalpha(c))
            f[0] ++;
        else if (isdigit(c))
            f[1] ++;
        else
            f[2] ++;
    }
    for (auto [k, c] : mp) {
        if (c > 2)
            ok = 0;
    }
    if (f[0] and f[1] and f[2]) {
        cout << (ok ? 2 : 1) << "\n";
    } else {
        cout << 0 << "\n";
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}