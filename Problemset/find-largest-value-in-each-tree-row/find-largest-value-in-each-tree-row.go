
// @Title: 在每个树行中找最大值 (Find Largest Value in Each Tree Row)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 22:17:17
// @Runtime: 8 ms
// @Memory: 5.5 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	result:=make([]int,0)
	stack:=[]*TreeNode{root}
	for len(stack)>0{
		length:=len(stack)
		maxVal:=-(1<<31)
		for i:=0;i<length;i++{
			node:=stack[0]
			stack=stack[1:]
			maxVal=max(maxVal,node.Val)
			if node.Left!=nil{
				stack=append(stack,node.Left)
			}
			if node.Right!=nil{
				stack=append(stack,node.Right)
			}
		}
		result=append(result,maxVal)
	}
	
	
	return result

}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
