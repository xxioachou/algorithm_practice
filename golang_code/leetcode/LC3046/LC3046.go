package main

func isPossibleToSplit(nums []int) bool {
    mp := make(map[int]int)
	for _, x := range nums {
		mp[x] ++
		if mp[x] > 2 {
			return false
		}
	}
	return true
}

func main() {

}