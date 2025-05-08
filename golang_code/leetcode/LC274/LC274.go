package main

func hIndex(citations []int) int {
    low, high := 0, len(citations)
	check := func(h int) bool {
		cnt := 0
		for _, v := range citations {
			if v >= h {
				cnt ++
			}
		}
		return cnt >= h
	}
	for low < high {
		mid := (low + high + 1) / 2
		if check(mid) {
			low = mid;
		} else {
			high = mid - 1
		}
	}
	return high
}

func main() {

}