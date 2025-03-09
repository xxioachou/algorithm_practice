package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := len(nums) / 2
	p := TreeNode{nums[idx], nil, nil}
	p.Left = sortedArrayToBST(nums[:idx])
	p.Right = sortedArrayToBST(nums[idx+1:])
	return &p
}

func main() {

}
