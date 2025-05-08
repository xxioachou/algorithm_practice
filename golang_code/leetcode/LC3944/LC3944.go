package main

func minimumCoins(prices []int) int {
	const inf = 0x3f3f3f3f
    n := len(prices)
	f := make([]int, n + 2)
	for i := range f {
		f[i] = inf
	}
	f[1] = 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n + 1 && j - i <= i + 1; j++ {
			f[j] = min(f[j], f[i] + prices[i - 1])
		}
	}
	return f[n + 1]
}

func main() {

}