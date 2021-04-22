
// @Title: 删除二叉搜索树中的节点 (Delete Node in a BST)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 15:38:40
// @Runtime: 16 ms
// @Memory: 6.9 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return nil 
	}
	if root.Val==key{
		if root.Left==nil{
			return root.Right
		}
		if root.Right==nil{
			return root.Left
		}
		root.Val=getMin(root.Right)
		root.Right=deleteNode(root.Right,root.Val)

	}else if root.Val>key{
		root.Left=deleteNode(root.Left,key)
	}else if root.Val<key{
		root.Right=deleteNode(root.Right,key)
	}
	return root
}

func getMin(node *TreeNode)int{
	for node.Left!=nil{
		node=node.Left
	}
	return node.Val
}
