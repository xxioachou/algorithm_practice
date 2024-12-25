package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	ok := true
	n := len(grid)
	for i := 0; i < n && ok; i++ {
		for j := 0; j < n && ok; j++ {
			if grid[i][j] != grid[0][0] {
				ok = false
			}
		}
	}
	if ok {
		val := false
		if grid[0][0] > 0 {
			val = true
		}
		return &Node{val, true, nil, nil, nil, nil}
	}

	p := Node{true, false, nil, nil, nil, nil}
	tl := make([][]int, n/2)
	tr := make([][]int, n/2)
	dl := make([][]int, n/2)
	dr := make([][]int, n/2)
	for i := 0; i < n/2; i++ {
		tl[i] = make([]int, n/2)
		tr[i] = make([]int, n/2)
		copy(tl[i], grid[i][:n/2])
		copy(tr[i], grid[i][n/2:])
	}
	for i := n / 2; i < n; i++ {
		dl[i-n/2] = make([]int, n/2)
		dr[i-n/2] = make([]int, n/2)
		copy(dl[i-n/2], grid[i][:n/2])
		copy(dr[i-n/2], grid[i][n/2:])
	}
	p.TopLeft, p.TopRight, p.BottomLeft, p.BottomRight =
		construct(tl), construct(tr), construct(dl), construct(dr)

	return &p
}

func main() {

}
