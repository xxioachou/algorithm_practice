package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root.Left) + count(root.Right) + 1
}

func kthSmallest(root *TreeNode, k int) int {
	c := count(root.Left)
	if c >= k {
		return kthSmallest(root.Left, k)
	} else if c+1 >= k {
		return root.Val
	}

	return kthSmallest(root.Right, k-c-1)
}

func main() {

}
