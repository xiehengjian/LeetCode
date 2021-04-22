
// @Title: 二叉树最大宽度 (Maximum Width of Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 08:28:24
// @Runtime: 0 ms
// @Memory: 3.7 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
    //. 1
    // 1. 2
    //1 2 3 4
 //1 2 3 4 5 6 7 8

 //
 if root==nil{
     return 0
 }
 root.Val=1
 maxWidth:=0
subarr:=make([]*TreeNode,0)
 queue:=[]*TreeNode{root}
 for len(queue)>0{
     length:=len(queue)
     subarr=make([]*TreeNode,0)
     for i:=0;i<length;i++{
         node:=queue[0]
         queue=queue[1:]
        subarr=append(subarr,node)
        if node.Left!=nil{
            node.Left.Val=2*node.Val-1
            queue=append(queue,node.Left)
        }
        if node.Right!=nil{
            node.Right.Val=2*node.Val
            queue=append(queue,node.Right)
        }
     }
     maxWidth=max(maxWidth,subarr[len(subarr)-1].Val-subarr[0].Val+1)

 }

 return maxWidth

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
