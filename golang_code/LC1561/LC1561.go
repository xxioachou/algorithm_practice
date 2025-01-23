package main

import "slices"

func maxCoins(piles []int) int {
	slices.Sort(piles)
	ans := 0
	n := len(piles) / 3
	for i := len(piles) - 1; i - 1 >= n; i -= 2 {
		ans += piles[i - 1];
	}
	return ans
}

func main() {

}