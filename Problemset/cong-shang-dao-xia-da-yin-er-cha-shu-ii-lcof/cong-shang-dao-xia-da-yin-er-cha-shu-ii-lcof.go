
// @Title: 从上到下打印二叉树 II (从上到下打印二叉树 II LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 14:16:31
// @Runtime: 0 ms
// @Memory: 2.8 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
ret:=[][]int{}//结果数组
	if root==nil{
		return ret 
	}
	queue:=[]*TreeNode{root}
	for len(queue)>0{
		subarr:=[]int{}
		length:=len(queue)
		for i:=0;i<length;i++{
			node:=queue[0]
			queue=queue[1:]
			subarr=append(subarr,node.Val)
			if node.Left!=nil{
				queue=append(queue,node.Left)
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
		}
		ret=append(ret,subarr)
	}
	return ret
}
