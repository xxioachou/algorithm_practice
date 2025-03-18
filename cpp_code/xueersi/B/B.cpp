#include <bits/stdc++.h>
using namespace std;

typedef long long ll;
const int N = 21010;

ll h[N], cnt[N];
ll ss[N], s[N];
bool ban[N];

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int n, m, x;
    cin >> n >> m >> x;
    for (int i = 1, h; i <= n; i ++) {
        cin >> h;
        cnt[h] ++;
    }
    for (int i = 1, p; i <= m; i ++) {
        cin >> p;
        ban[h[p]] = true;
    }
    for (int i = 1; i <= 20000; i ++) {
        ss[i] = ss[i - 1];
        s[i] = s[i - 1];
        if (cnt[i]) {
            ss[i] += i * i;
            s[i] += i;
        }
        cnt[i] += cnt[i - 1];
    }
    for (int i = 1; i <= 20000; i ++) {

        ll cost = 0;
        if (i + x <= 20000) {
            int l = i + x;
            ll len = cnt[20000] - cnt[l];
            cost += len * (ss[20000] - ss[l]);
            cost -= 2ll * len * l * (s[20000] - s[l]);
            cost += l * l * len;
        }
        if (i - x)

        if (ban[i]) {
            break;
        }
    }


    return 0;
}