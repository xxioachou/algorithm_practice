package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	p := TreeNode{preorder[0], nil, nil}
	i := 0
	for ; inorder[i] != p.Val; i++ {
	}

	p.Left = buildTree(preorder[1:], inorder[:i])
	p.Right = buildTree(preorder[1+i:], inorder[i+1:])
	return &p
}

func main() {

}
