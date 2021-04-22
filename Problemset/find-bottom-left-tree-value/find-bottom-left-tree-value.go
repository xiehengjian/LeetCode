
// @Title: 找树左下角的值 (Find Bottom Left Tree Value)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 22:43:08
// @Runtime: 8 ms
// @Memory: 6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	
	stack:=[]*TreeNode{root}
	result:=make([][]int,0)
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
	
	return result[len(result)-1][0]

}
