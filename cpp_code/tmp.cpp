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

void solve() {

}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}