#include <bits/stdc++.h>
using namespace std;

class ExamRoom {
public:
    int N;
    // 最大距离，位置
    set<pair<int, int>> st;
    set<int> nums;
    ExamRoom(int n) {
        N = n;
    }
    
    void insert(int l, int r) {
        if (l + 1 < r) {
            int mid = (l + r) / 2;
            int len = mid - l;
            st.emplace(-len, mid);
        }
    }

    void erase(int l, int r) {
        if (l + 1 < r) {
            int mid = (l + r) / 2;
            int len = mid - l;
            st.erase(st.find({-len, mid}));
        }
    }

    int seat() {
        if (nums.empty()) {
            nums.emplace(0);
            return 0;
        } 

        int dis = 0, p = -1;
        if (dis < *nums.begin() - 0) {
            dis = *nums.begin() - 0;
            p = 0;
        }
        if (!st.empty()) {
            auto [d, pos] = *st.begin();
            if (dis < -d) {
                dis = -d;
                p = pos;
            }
        }
        if (dis < N - 1 - *prev(nums.end())) {
            dis = N - 1 - *prev(nums.end());
            p = N - 1;
        }
        if (!st.empty() && p == st.begin()->second) {
            st.erase(st.begin());
        }

        auto it = nums.upper_bound(p);
        if (it != nums.end()) {
            int r = *it;
            insert(p, r);
        }

        if (it != nums.begin()) {
            int l = *prev(it);
            insert(l, p);
        }

        nums.emplace(p);
        return p;
    }
    
    void leave(int p) {
        auto it = nums.lower_bound(p);
        auto it2 = next(it);
        if (it != nums.begin() && it2 != nums.end()) {
            int l = *prev(it);
            int r = *it2;
            erase(l, p);
            erase(p, r);
            insert(l, r);
        }
        nums.erase(it);   
    }
};

void solve() {
    int N;
    cin >> N;
    ExamRoom e(N);
    for (int i = 0; i < N; i++) {
        cout << "seat: " << e.seat() << '\n';
    }
    srand((unsigned)time(nullptr));
    vector<int> vis(N);
    for (int i = 0; i < 4; i++) {
        e.leave(i);
    }
    for (int i = 6; i < N; i++) {
        e.leave(i);
    }
    cout << e.seat() << '\n';
    cout << e.seat() << '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr), cout.tie(nullptr);
    int t = 1;
    // cin >> t;
    while (t--) solve();
    return 0;
}