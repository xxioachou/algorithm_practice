#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

/*

    R G B Y W -> 5 6 7 8 9
    1 2 3 4 5 -> 0 1 2 3 4
*/

void solve() {
    int n;
    cin >> n;  
    vector<int> col(n), val(n);
    vector<vector<int>> st(10); 
    string s;
    for (int i = 0; i < n; i++) {
        cin >> s;
        if (s[0] == 'R') {
            col[i] = 5;
        } else if (s[0] == 'G') {
            col[i] = 6;
        } else if (s[0] == 'B') {
            col[i] = 7;
        } else if (s[0] == 'Y') {
            col[i] = 8;
        } else if (s[0] == 'W') {
            col[i] = 9;
        }
        val[i] = s[1] - '1';
        st[col[i]].push_back(i);
        st[val[i]].push_back(i);
    }
    int ans = 10;
    vector<int> vis(n);
    for (int i = 0; i < (1 << 10); i++) {
        int cnt = 0;
        fill(vis.begin(), vis.end(), 0);
        for (int j = 0; j < n; j++) {
            if ((i >> col[j] & 1) and (i >> val[j] & 1)) {
                // this card is distinct
                vis[j] = 1;
                ++ cnt;
            }
        }
        while (true) {
            int x = cnt;
            for (int j = 0; j < n; j++) {
                if (vis[j])
                    continue;
                int value = -1, ok1 = 1;
                if (i >> col[j] & 1) {
                    for (int x : st[col[j]]) {
                        if (vis[x])
                            continue;
                        if (value == -1)
                            value = val[x];
                        else if (value != val[x]) {
                            ok1 = 0;
                            break;
                        }
                    }
                } else {
                    ok1 = 0;
                }
                int color = -1, ok2 = 1;
                if (i >> val[j] & 1) {
                    for (int x : st[val[j]]) {
                        if (vis[x])
                            continue;
                        if (color == -1)
                            color = col[x];
                        else if (color != col[x]) {
                            ok2 = 0;
                            break;
                        }
                    }
                } else {
                    ok2 = 0;
                }
                if (ok1 or ok2) {
                    vis[j] = 1;
                    ++ cnt;
                }
            }
            if (x == cnt)
                break;
        }
  
        if (cnt >= n) {    
            // all cards are distinct
            ans = min(ans, __builtin_popcount(i));
        } else {        
            // only one type of card left
            int lc = -1, lv = -1, ok = 1;
            for (int j = 0; j < n; j++) {
                if (vis[j])
                    continue;
                if (lc == -1) {
                    lc = col[j];
                    lv = val[j];
                } else if (lc != col[j] or lv != val[j]) {
                    ok = 0;
                    break;
                }
            }
            if (ok) {
                ans = min(ans, __builtin_popcount(i));
            }
        }
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