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
    int cnt = 0;
    for (int i = 0, j = 1; i < s.size(); ) {
        j = i + 1;
        while (j < s.size() && s[j] == s[i]) {
            j ++;
        }
        cnt ++;
        i = j;
    }
    cout << (cnt <= 5 ? "Yes\n" : "No\n");
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}