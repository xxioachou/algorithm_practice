package main

func numberOfAlternatingGroups(colors []int, k int) int {
	ans, n := 0, len(colors)
	for i := 0; i < n+k-1; {
		j := i + 1
		for j < n+k-1 && colors[j%n] != colors[(j-1)%n] {
			j++
		}
		ans += max(j-i-k+1, 0)
		i = j
	}
	return ans
}

func main() {

}
