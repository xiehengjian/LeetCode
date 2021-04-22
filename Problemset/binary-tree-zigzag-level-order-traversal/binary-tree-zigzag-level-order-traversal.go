
// @Title: 二叉树的锯齿形层序遍历 (Binary Tree Zigzag Level Order Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 19:04:50
// @Runtime: 0 ms
// @Memory: 2.5 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret:=[][]int{}//结果数组
	flag:=true
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
		if !flag{
			subarr=reverse(subarr)
		}
		flag=!flag
		ret=append(ret,subarr)
	}
	return ret

}

func reverse(a []int)[]int{
	b:=make([]int,0)
	for i:=len(a)-1;i>-1;i--{
		b=append(b,a[i])
	}
	return b

}
