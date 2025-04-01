#include <bits/stdc++.h>
#define x first
#define y second
#define sz(x) ((int)x.size())
using namespace std;

typedef long long ll;
const int inf = 0x3f3f3f3f;
const ll INF = 1e18;

void solve() {
	int n, a;
	cin >> n >> a;
	vector<int> x(n);
	for (auto &e : x) {
		cin >> e;
	}
	sort(x.begin(), x.end());
	int ans = 0;
	if (sz(x) == 1) {
		
	} else {
		if (a <= x[0]) {
			ans = x[n - 2] - a;
		} else if (a >= x[n - 1]) {
			ans = a - x[1];
		} else {
			ans = min(abs(a - x[0]), abs(a - x[n - 2])) + abs(x[n - 2] - x[0]);
			ans = min(ans, min(abs(x[n - 1] - a), abs(a - x[1])) + x[n - 1] - x[1]);
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