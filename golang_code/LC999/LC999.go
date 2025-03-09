package main

func numRookCaptures(board [][]byte) int {
    x, y := 0, 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				x, y = i, j
			}
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	ans := 0
	for i := 0; i < 4; i++ {		
		for tx, ty := x, y; tx >= 0 && tx < 8 && ty >= 0 && ty < 8 &&
			board[tx][ty] != 'B'; tx, ty = tx + dx[i], ty + dy[i] {
			if board[tx][ty] == 'p' {
				ans ++
                break
			}
		}
	}
	return ans
} 

func main() {

}