
// @Title: 二叉树的层序遍历 II (Binary Tree Level Order Traversal II)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 16:57:19
// @Runtime: 4 ms
// @Memory: 2.9 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	result:=make([][]int,0)
	stack:=[]*TreeNode{root}
	for len(stack)>0{
		length:=len(stack)
		subarr:=make([]int,0)
		for i:=0;i<length;i++{
			node:=stack[0]
			stack=stack[1:]
			subarr=append(subarr,node.Val)
			if node.Left!=nil{
				stack=append(stack,node.Left)
			}
			if node.Right!=nil{
				stack=append(stack,node.Right)
			}
		}
		result=append(result,subarr)
	}
	res:=make([][]int,0)
	for len(result)>0{
		res=append(res,result[len(result)-1])
		result=result[:len(result)-1]
	}
	return res

}
