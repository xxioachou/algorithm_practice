#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
const int mod1 = 1e9 + 7;
const int mod2 = 1e9 + 181;

ll qpow(ll a, ll b, ll mod) {
    ll res = 1;
    while (b) {
        if (b & 1)
            res = res * a % mod;
        b >>= 1;
        a = a * a % mod;
    }
    return res % mod;
}

void solve() {
    int n, p;
    cin >> n >> p;
    vector<int> k(n);
    for (auto &x : k) {
        cin >> x;
    }    
    ll s1 = 0, s2 = 0;
    sort(k.rbegin(), k.rend());
    for (auto &x : k) {
        if (s1 == 0 and s2 == 0) {      
            // real zero
            s1 = (s1 + qpow(p, x, mod1)) % mod1;
            s2 = (s2 + qpow(p, x, mod2)) % mod2;
        } else {
            // greater than zero
            s1 = (s1 - qpow(p, x, mod1) + mod1) % mod1;
            s2 = (s2 - qpow(p, x, mod2) + mod2) % mod2;
        }
    }
    cout << s1 << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    cin >> t;
    while (t--) solve();
    return 0;
}