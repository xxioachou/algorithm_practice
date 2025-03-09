#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
using st = bitset<510>;

void solve() {
    ll N, X, K;
    cin >> N >> X >> K;
    vector<ll> P(N + 1), U(N + 1), C(N + 1);
    for (int i = 1; i <= N; i++) {
        cin >> P[i] >> U[i] >> C[i];
    }
    vector<vector<map<st, ll>>> f(N + 1, vector<map<st, ll>>(X + 1));
    f[0][0][st(0)] = 0;
    for (int i = 0; i < N; i++) {
        for (int j = 0; j <= X; j++) {
            if (j + P[i] <= X) {
                for (auto [s, v] : f[i][j]) {
                    st tmp;
                    tmp[C[i]] = 1;
                    auto state = s | tmp;
                    ll val = -K * s.count() + K * state.count();
                    f[i + 1][j + P[i]][state] = max(f[i + 1][j + P[i]][state],
                    v + val + U[i]);
                }
            }
        }
    }
    ll ans = 0;
    for (int j = 0; j <= X; j++) {
        for (auto &[s, v] : f[N][j]) {
            ans = max(ans, v);
        }
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