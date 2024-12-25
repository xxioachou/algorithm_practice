package main

func knightDialer(n int) int {
    f := make([][]int, 2)
	f[0] = make([]int, 10)
	f[1] = make([]int, 10)
	e := make([][]int, 10)
	e[0] = append(e[0], 4, 6)
	e[1] = append(e[1], 6, 8)
	e[2] = append(e[2], 7, 9)
	e[3] = append(e[3], 4, 8)
	e[4] = append(e[4], 0, 3, 9)
	e[6] = append(e[6], 0, 1, 7)
	e[7] = append(e[7], 2, 6)
	e[8] = append(e[8], 1, 3)
	e[9] = append(e[9], 2, 4)
	for i := range(e) {
		f[1][i] = 1
	}
	const mod = 1000000007
	for i := 1; i < n; i++ {
		for j := range(e) {
			f[(i + 1) & 1][j] = 0
		}
		for j := range(e) {
			for _, k := range(e[j]) {
				f[(i + 1) & 1][k] += f[i & 1][j]
				f[(i + 1) & 1][k] %= mod
			}
		}
	}
	ans := 0
	for i := range(e) {
		ans += f[n & 1][i]
		ans %= mod
	}
	return ans
} 

func main() {

}