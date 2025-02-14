package main

func majorityElement(nums []int) int {
    ans, count := 0, -1
	for _, v := range nums {
		if count == 0 {
			ans = v
		}
		if ans != v {
			count ++
		} else {
			count --
		}
	}
	return ans
}

func main() {

}