package main

import "math"

func maxProfit(prices []int) int {
    minp := math.MaxInt
	ans := 0
	for _, price := range prices {
		ans = max(ans, price - minp)
		minp = min(minp, price)
	}
	return ans
}

func main() {

}