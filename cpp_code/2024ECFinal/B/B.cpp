#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

map<string, int> mp = {
    {"I", 1},
    {"II", 2},
    {"III", 3},
    {"IV", 4},
    {"V", 5},
    {"VI", 6},
    {"VII", 7},
    {"VIII", 8},
};
void solve() {
    string s;
    cin >> s;
    int n = s.size();
    s = " " + s;
    vector<int> g(n + 2, inf);
    g[n + 1] = 0;
    for (int i = n; i; i--) {
        for (int j = 1; j <= 4 && i + j - 1 <= n; j++) {
            string str = s.substr(i, j);
            if (mp.count(str))
                g[i] = min(g[i], g[i + j] + 1);
        }  
        // cout << ":" << i << ' ' << g[i] << '\n';      
    }
    int m = g[1];
    for (int i = 1; i <= n; ) {
        string str = "";
        for (int j = 1; j <= 4 && i + j - 1 <= n; j++) {
            str.push_back(s[i + j - 1]);
            if (mp.count(str) && g[i + j] == m - 1) {
                cout << mp[str];
                i += j;
                m --;
                break;
            }
        }
    }
    cout << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}