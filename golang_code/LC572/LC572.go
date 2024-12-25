package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func same(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	}
	if q == nil {
		return p == nil
	}
	return p.Val == q.Val && same(p.Left, q.Left) && same(p.Right, q.Right)
}

func get_height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(get_height(root.Left), get_height(root.Right)) + 1
}

func dfs(root *TreeNode, subRoot *TreeNode, subHeight *int, ok *bool) int {
	if root == nil {
		return 0
	}

	height := max(dfs(root.Left, subRoot, subHeight, ok), dfs(root.Right, subRoot, subHeight, ok)) + 1
	if height == *subHeight {
		if same(root, subRoot) {
			fmt.Println(root.Val)
			*ok = true
		}
	}
	return height
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	subHeight := get_height(subRoot)
	ok := false
	// fmt.Println(ok, subHeight)
	dfs(root, subRoot, &subHeight, &ok)
	return ok
}

func main() {

}
