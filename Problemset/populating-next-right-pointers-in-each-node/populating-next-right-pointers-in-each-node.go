
// @Title: 填充每个节点的下一个右侧节点指针 (Populating Next Right Pointers in Each Node)
// @Author: 846188037@qq.com
// @Date: 2021-04-10 16:23:37
// @Runtime: 4 ms
// @Memory: 6.3 MB

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root==nil{
		return root
	}
	connectNodes(root.Left,root.Right)
	return root
}

func connectNodes(node1 *Node,node2 *Node){
	if node1==nil || node2==nil{
		return
	}
	node1.Next=node2 
	connectNodes(node1.Left,node1.Right)
	connectNodes(node2.Left,node2.Right)
	connectNodes(node1.Right,node2.Left)
}
