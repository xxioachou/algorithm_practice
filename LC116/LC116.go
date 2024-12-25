package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	var dfs func(*Node, *Node)
	dfs = func(p *Node, q *Node) {
		if p == nil {
			return
		}
		p.Next = q
		if q != nil {
			dfs(p.Right, q.Left)
		} else {
			dfs(p.Right, nil)
		}
		dfs(p.Left, p.Right)
	}
	dfs(root, nil)
	return root
}

func main() {

}
