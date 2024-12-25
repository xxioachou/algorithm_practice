#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
    int n = 7000;
    ofstream fout("6.in");
    for (int i = 1; i <= n; i++) {
        fout << (1ll << 60) - 1 << ' ';
    }
    fout << "\n";
    for (int i = 1; i <= n; i++) {
        fout << 1 << ' ';
    }
    fout << '\n';
    fout.close();
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}