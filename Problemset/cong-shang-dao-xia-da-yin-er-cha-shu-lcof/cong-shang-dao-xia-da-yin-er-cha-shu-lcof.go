
// @Title: 从上到下打印二叉树 (从上到下打印二叉树 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 14:15:33
// @Runtime: 0 ms
// @Memory: 2.6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
    if root==nil{
        return []int{}
    }
    queue:=[]*TreeNode{root}
    res:=[]int{}
    for len(queue)>0{
        node:=queue[0]
        queue=queue[1:]
        res=append(res,node.Val)
        if node.Left!=nil{
            queue=append(queue,node.Left)
        }
        if node.Right!=nil{
            queue=append(queue,node.Right)
        }
    }
    return res

}
