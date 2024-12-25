package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		res[i] = make([]int, len(img[i]))
	}
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			cnt, sum := 0, 0
			for k := -1; k <= 1; k++ {
				for u := -1; u <= 1; u++ {
					if i+k < 0 || i+k >= len(img) ||
						j+u < 0 || j+u >= len(img[i]) {
						continue
					}
					cnt++
					sum += img[i+k][j+u]
				}
			}
			fmt.Println(i, j, cnt, sum)
			res[i][j] = sum / cnt
		}
	}
	return res
}

func main() {

}
