#include <bits/stdc++.h>
using namespace std;

/*
    随机底数 a，固定模数 p = 1e9 + 7 的双哈希
    s 是长度为 n 的字符串
    h = (a^(n-1) * s[0] + a^(n - 2) * s[1] + ... + a^0 * s[n - 1]) % p
 */
seed_seq seq {
    (uint64_t) chrono::duration_cast<chrono::nanoseconds>(chrono::steady_clock::now().time_since_epoch()).count(),
    (uint64_t) __builtin_ia32_rdtsc(),
    (uint64_t) (uintptr_t) make_unique<char>().get()
};
mt19937 rng(seq);
const int p = 1e9 + 7;
const int a1 = uniform_int_distribution<int>(0, p - 1)(rng);
const int a2 = uniform_int_distribution<int>(0, p - 1)(rng);
const int N = 400010;

int h1[N], h2[N], p1[N], p2[N];

pair<int, int> get_hash(int l, int r) {
    int x = h1[r] - (1ll * h1[l - 1] * p1[r - l + 1] % p);
    int y = h2[r] - (1ll * h2[l - 1] * p2[r - l + 1] % p);
    x = (x % p + p) % p;
    y = (y % p + p) % p;
    return pair<int, int>{x, y};
}

void solve() {
    string s;
    cin >> s;
    int n = s.size();
    p1[0] = p2[0] = 1;
    for (int i = 0; i < n; i++) {
        p1[i + 1] = 1ll * p1[i] * a1 % p;
        p2[i + 1] = 1ll * p2[i] * a2 % p;
        h1[i + 1] = (1ll * h1[i] * a1 + s[i] - 'a') % p;
        h2[i + 1] = (1ll * h2[i] * a2 + s[i] - 'a') % p;
    }
    for (int t = 1; n - t * 2 > 0; t++) {
        int x = n - t * 2;
        int len = t + x;
        if (get_hash(n - t + 1, n) != get_hash(n - t*2 + 1, n - t))
            continue;
        if (get_hash(1, x) != get_hash(n - x + 1, n))
            continue;
        cout << "YES\n" << s.substr(0, len) << '\n';
        return;
    }
    cout << "NO\n";
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}