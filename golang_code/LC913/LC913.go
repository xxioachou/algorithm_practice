package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "time"
)

func catMouseGame(graph [][]int) int {
	const CAT_TURN = 1
	const MOUSE_TURN = 0
	const UNKNOWN = 0
	const MOUSE_WIN = 1
	const CAT_WIN = 2
	const HOLE = 0

	n := len(graph)
	getPrevState := func(turn, mouse, cat int) [][]int {
		states := make([][]int, 0)
		if turn == MOUSE_TURN {
			for _, x := range graph[cat] {
				if x != HOLE {
					states = append(states, []int{CAT_TURN, mouse, x})
				}
			}
		} else {
			for _, x := range graph[mouse] {
				states = append(states, []int{MOUSE_TURN, x, cat})
			}
		}
		return states
	}

	f := make([][][]int, 2)
	deg := make([][][]int, 2)
	for i := range f {
		f[i] = make([][]int, n)
		deg[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
			deg[i][j] = make([]int, n)
		}
	}
	que := make([][]int, 0)
	for j := 1; j < n; j ++ {
		f[MOUSE_TURN][HOLE][j] = MOUSE_WIN
		f[CAT_TURN][HOLE][j] = MOUSE_WIN
		que = append(que, []int{MOUSE_TURN, HOLE, j})
		que = append(que, []int{CAT_TURN, HOLE, j})
	}
	for j := 1; j < n; j ++ {
		f[MOUSE_TURN][j][j] = CAT_WIN
		f[CAT_TURN][j][j] = CAT_WIN
		que = append(que, []int{MOUSE_TURN, j, j})
		que = append(que, []int{CAT_TURN, j, j})
	}
	for j := 0; j < n; j ++ {
		for k := 1; k < n; k ++ {
			deg[MOUSE_TURN][j][k] = len(graph[j])
			deg[CAT_TURN][j][k] = len(graph[k])
		}
	}
	// 猫不能走到 0
	for _, v := range graph[HOLE] {
		for i := 0; i < n; i ++ {
			deg[CAT_TURN][i][v] --
		}
	}

	for len(que) > 0 {
		stat := que[0]
		que = que[1:]
		turn, mouse, cat := stat[0], stat[1], stat[2]
		prev_states := getPrevState(turn, mouse, cat)
		winner := f[turn][mouse][cat]
		for _, prev_stat := range prev_states {
			prev_turn, prev_mouse, prev_cat := prev_stat[0], prev_stat[1], prev_stat[2]
			if f[prev_turn][prev_mouse][prev_cat] == UNKNOWN {
				win := (prev_turn == MOUSE_TURN && winner == MOUSE_WIN) || 
				(prev_turn == CAT_TURN && winner == CAT_WIN)
				if win {
					f[prev_turn][prev_mouse][prev_cat] = winner
					que = append(que, []int{prev_turn, prev_mouse, prev_cat})
				} else {
					deg[prev_turn][prev_mouse][prev_cat] --
					if deg[prev_turn][prev_mouse][prev_cat] == 0 {
						f[prev_turn][prev_mouse][prev_cat] = winner
						que = append(que, []int{prev_turn, prev_mouse, prev_cat})
					}	
				}
			}
		}
	}
	return f[MOUSE_TURN][1][2]
}

func main() {
	var n int
	fmt.Scan(&n)
	g := make([][]int, n)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		str, _ := in.ReadString('\n')
		parts := strings.Fields(str)
		g[i] = make([]int, len(parts))
		for j := range parts {
			g[i][j], _ = strconv.Atoi(parts[j])
		}
	}
	// fmt.Println(n, g)
	fmt.Println(catMouseGame(g))
}