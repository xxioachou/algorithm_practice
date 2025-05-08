package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	p := TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for ; inorder[i] != p.Val; i++ {
	}
	p.Right = buildTree(inorder[i+1:], postorder[:len(postorder)-1])
	p.Left = buildTree(inorder[:i], postorder[:len(postorder)-(len(inorder)-i)])
	return &p
}

func main() {

}
