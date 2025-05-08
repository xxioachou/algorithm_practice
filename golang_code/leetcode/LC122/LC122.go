package main

func maxProfit(prices []int) int {
	const inf = 0x3f3f3f3f
    f0, f1 := 0, -inf
	for _, p := range prices {
		nf0, nf1 := f0, f1
		nf1 = max(nf1, f0 - p)
		nf0 = max(nf0, f1 + p)
		f0, f1 = nf0, nf1
	}
	return max(f0, f1)
}

func main() {

}