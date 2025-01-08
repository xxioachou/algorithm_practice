package main

import (
	"math"
	"strconv"
)

func largestGoodInteger(num string) string {
    maxv := math.MinInt
	ans := ""
	for i := 0; i + 2 < len(num); i++ {
		if num[i] == num[i + 1] && num[i] == num[i + 2] {
			tmp, _ := strconv.Atoi(num[i : i + 3])
			if maxv < tmp {
				maxv = tmp
				ans = num[i : i + 3]
			}
		}
	}
	return ans

}

func main() {

}
