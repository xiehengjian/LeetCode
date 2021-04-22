
// @Title: 二叉树的前序遍历 (Binary Tree Preorder Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 14:20:01
// @Runtime: 0 ms
// @Memory: 2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root==nil{
		return nil 
	}
	stack:=[]*TreeNode{root}
	result:=[]int{}
	for len(stack)!=0{
		node:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		result=append(result,node.Val)
		if node.Right!=nil{
			stack=append(stack,node.Right)
		}
		if node.Left!=nil{
			stack=append(stack,node.Left)
		}

	}
	return result

}
