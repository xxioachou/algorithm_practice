package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	index 	int
	inorder []int
}

func Constructor(root *TreeNode) BSTIterator {
	var obj BSTIterator
	
	obj.inorder = append(obj.inorder, -1)
	obj.index = 0

    var dfs func(*TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Left)
		obj.inorder = append(obj.inorder, cur.Val)
		dfs(cur.Right)
	}
	dfs(root)
	return obj
}


func (this *BSTIterator) Next() int {
	this.index = this.index + 1
    return this.inorder[this.index]	
}


func (this *BSTIterator) HasNext() bool {
    return this.index + 1 < len(this.inorder)
} 

func main() {

}