#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
	int n;
	cin >> n;
	vector<int> a(n);
	vector<vector<pair<int, int>>> p(n + 1);
	for (auto &x : a) {
		cin >> x;
	}
	for (int i = 0; i < n; ) {
		int j = i + 1;
		while (j < n && a[j] == a[i]) {
			j ++;
		}
		p[a[i]].emplace_back(i, j - 1);
		i = j;
	}
	int ans = n;
	for (int i = 1; i <= n; i ++) {
		if (p[i].empty()) {
			continue;
		}
		int l = -1, r = n;
		for (auto [x, y] : p[i]) {
			if (x == l + 1) {
				l = y;
			} else {
				break;
			}
		}
		l ++;

		for (auto it = p[i].rbegin(); it != p[i].rend(); it ++) {
			auto [x, y]	= *it;
			if (y == r - 1) {
				r = x;
			} else {
				break;
			}
		}
		r --;
		
		int res = max(r - l + 1, 0);
		ans = min(ans, res);
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