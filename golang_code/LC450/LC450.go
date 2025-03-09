package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val != key {
		if key < root.Val {
			root.Left = deleteNode(root.Left, key)
		} else {
			root.Right = deleteNode(root.Right, key)
		}
		return root
	}

	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}

	ans := root.Right
	p := root.Right
	for p.Left != nil {
		p = p.Left
	}
	p.Left = root.Left
	return ans
}

func main() {

}
