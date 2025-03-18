#include <bits/stdc++.h>
#include <cmath>
using namespace std;

typedef long long ll;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int N, cur = 0, M;
    cin >> N;
    vector<ll> a(N);
    for (auto &x : a) {
        cin >> x;
    }
    cin >> M;
    while (M-- > 0) {
        ll P, sum = 0;
        cin >> P;
        for (auto &x : a) {
            sum += x * x;
        }
        P -= (P / sum) * sum;
        while (P > 0) {
            int ne = (cur + 1) % N;
            P -= a[ne] * a[ne];
            cur = ne;
        }
        cout << cur + 1 << '\n';
        a[cur] ++;
        // cout << ":";
        // for (auto &x : a) {
        //     cout << x << ' ';
        // }
        // cout << '\n';
    }
    return 0;
}