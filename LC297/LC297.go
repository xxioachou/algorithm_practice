package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

func dfs1(cur *TreeNode) string {
	if cur == nil {
		return "*"
	}
	return strconv.Itoa(cur.Val) + "," + dfs1(cur.Left) + "," + dfs1(cur.Right)
}

func dfs2(a []string, index *int) *TreeNode {
	if *index >= len(a) || a[*index] == "*" {
		return nil
	}

	val, _ := strconv.Atoi(a[*index])
	*index += 1

	r := TreeNode{val, nil, nil}
	r.Left = dfs2(a, index)
	r.Right = dfs2(a, index)

	return &r
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return dfs1(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	index := 0
	return dfs2(strings.Split(data, ","), &index)
}
 

func main() {

}