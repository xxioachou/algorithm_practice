package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var head, last, cur, pre *Node
	pre = root
	// 构建完上一层的节点，根据上一层的信息构建当前这层节点
	for {
		// 找当前层第一个子结点
		for pre != nil {
			if pre.Left != nil {
				head = pre.Left
			} else {
				head = pre.Right
			}
			if head != nil {
				if head == pre.Right {
					pre = pre.Next
				}
				break
			}
			pre = pre.Next
		}
		// 没有下一层了
		if head == nil {
			break
		}

		last = head
		// 构建当前的层
		for last != nil {
			cur = nil
			for pre != nil {
				if pre.Left != nil && pre.Left.Next == nil {
					cur = pre.Left
					break
				}
				if pre.Right != nil {
					cur = pre.Right
					pre = pre.Next
					break
				}
				pre = pre.Next
			}
			last.Next = cur
			last = cur
		}
		pre = head
	}
	return root
}

func main() {

}
