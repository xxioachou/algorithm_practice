package main

func largestCombination(candidates []int) int {
	ans := 0
    for i := 29; i >= 0; i-- {
		tmp := 0
		for _, x := range candidates {
			if (x >> i & 1) != 0 {
				tmp ++
			}
		}
		ans = max(ans, tmp)
	}
	return ans
}

func main() {

}