
// @Title: N 叉树的层序遍历 (N-ary Tree Level Order Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 00:16:28
// @Runtime: 4 ms
// @Memory: 4.3 MB

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	ret:=[][]int{}//结果数组
	if root==nil{
		return ret 
	}
	queue:=[]*Node{root}
	for len(queue)>0{
		subarr:=[]int{}
		length:=len(queue)
		for i:=0;i<length;i++{
			node:=queue[0]
			queue=queue[1:]
			subarr=append(subarr,node.Val)
			for _,v:=range node.Children{
				queue=append(queue,v)
			}
		}
		ret=append(ret,subarr)
	}
	return ret
}
