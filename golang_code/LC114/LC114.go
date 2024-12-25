package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var dfs func(*TreeNode) []*TreeNode
	dfs = func(u *TreeNode) []*TreeNode {
		if u == nil {
			return []*TreeNode{nil, nil}
		}
		if u.Left == nil && u.Right == nil {
			return []*TreeNode{u, u}
		}

		if u.Left == nil {
			b := dfs(u.Right)
			u.Right = b[0]
			return []*TreeNode{u, b[1]}
		}
		if u.Right == nil {
			a := dfs(u.Left)
			u.Left = nil
			u.Right = a[0]
			return []*TreeNode{u, a[1]}
		}

		a := dfs(u.Left)
		b := dfs(u.Right)
		u.Left = nil
		u.Right = a[0]
		a[1].Right = b[0]
		// fmt.Println(u.Val, b[1].Val)
		return []*TreeNode{u, b[1]}
	}
	dfs(root)
}

func main() {

}
