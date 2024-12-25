#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;
constexpr int LEN = 80;
constexpr int MAXN = 100010;
constexpr int MAXM = 100010;
constexpr int CNT = MAXM / LEN + 10;

int N, M, Q;
// 将每次操作后的数列置于可持久化线段树上
int root[MAXM * 2], idx;
struct node {
    ll sum, v;
} tr[MAXN * 4 + MAXM * 17];

// 将询问分块
ll maxv[CNT], mins[CNT], maxs[CNT];

int insert(int p, int l, int r, int i, int x) {

}

void solve() {
    cin >> N >> M;
    for (int i = 1, L, R, X, t; i <= M; i++) {
        cin >> L >> R >> X;
        t = insert(root[i - 1], 1, N + 1, L, X);
        root[i] = insert(t, 1, N + 1, R + 1, X);
    }
    // 将询问分块
    for (int i = 1; i <= M; ) {
        int j = i;
        while (j <= M && j - i + 1 <= LEN) {

            j ++;
        }
        i = j;
    }
    cin >> Q;
    for (int i = 1, K, S, T; i <= Q; i++) {
        cin >> K >> S >> T;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}