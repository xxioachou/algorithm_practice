package main

func numberOfAlternatingGroups(colors []int) int {
	ans := 0
	for i := 0; i < len(colors); i++ {
		j := (i + 1) % len(colors)
		k := (i + 2) % len(colors)
		if (colors[j]^colors[i]) == 1 &&
			(colors[j]^colors[k]) == 1 {
			ans++
		}
	}
	return ans
}

func main() {

}
