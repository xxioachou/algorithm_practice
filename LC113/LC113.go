package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(*TreeNode, int)
	ans := make([][]int, 0)
	path := make([]int, 0)
	dfs = func(u *TreeNode, t int) {
		if u == nil {
			return
		}
		path = append(path, u.Val)
		// fmt.Println(u.Val, path)
		if u.Left == nil && u.Right == nil {
			if u.Val == t {
				tmp := make([]int, len(path))
				copy(tmp, path)
				ans = append(ans, tmp)
			}
			path = path[:len(path)-1]
			return
		}
		dfs(u.Left, t-u.Val)
		dfs(u.Right, t-u.Val)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return ans
}

func main() {

}
