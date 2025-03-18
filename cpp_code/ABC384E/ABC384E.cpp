#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

void solve() {
	ll H, W, X, P, Q;
	cin >> H >> W >> X >> P >> Q;
	P --, Q --;
	vector<vector<ll>> S(H, vector<ll>(W));
	vector<vector<int>> vis(H, vector<int>(W, 0));
	for (auto &v : S) {
		for (auto &x : v) {
			cin >> x;
		}
	}
	priority_queue<tuple<ll, int, int>, vector<tuple<ll, int, int>>, greater<tuple<ll, int, int>>> q;
	auto put = [&](int x, int y) -> void {
		for (int i = 0; i < 4; i ++) {
			int a = x + dx[i];
			int b = y + dy[i];
			if (a < 0 || a >= H || b < 0 || b >= W || vis[a][b]) {
				continue;
			}
			vis[a][b] = 1;
			q.emplace(S[a][b], a, b);
		}
	};
	ll ans = S[P][Q];
	vis[P][Q] = 1;
	put(P, Q);
	while (!q.empty()) {
		auto [v, i, j] = q.top();
		q.pop();

		if (v < (ans + X - 1) / X) {
			ans += v;
			put(i, j);
			// cout << ": " << i + 1 << ' ' << j + 1 << ' ' << ans << '\n';
		} else {
			break;
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