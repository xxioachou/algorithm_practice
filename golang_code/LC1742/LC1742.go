package main

func countBalls(lowLimit int, highLimit int) int {
	helper := func(x int) int {
		res := 0
		for x > 0 {
			res += x % 10
			x /= 10
		}
		return res
	}
	maxv := 0
	mp := make([]int, 50)
	for i := lowLimit; i <= highLimit; i ++ {
		idx := helper(i)
		mp[idx] ++
		if mp[idx] > maxv {
			maxv = mp[idx]
		}
	}
	return maxv
}

func main() {

}