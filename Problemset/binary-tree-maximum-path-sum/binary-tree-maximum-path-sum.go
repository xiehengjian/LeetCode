
// @Title: 二叉树中的最大路径和 (Binary Tree Maximum Path Sum)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 20:32:25
// @Runtime: 24 ms
// @Memory: 7.8 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Result struct{
    SingleMax int 
    Max int
}
func maxPathSum(root *TreeNode) int {
    if root==nil{
        return 0
    }
    res:=helper(root)
    return res.Max
}


func helper(root *TreeNode)Result{
    if root==nil{
        return Result{
            SingleMax:0,
            Max:-math.MaxInt32,
        }
    }
    left:=helper(root.Left)
    right:=helper(root.Right)
    res:=Result{SingleMax:0,Max:0,}
    res.SingleMax=max(left.SingleMax+root.Val,res.SingleMax)
    res.SingleMax=max(right.SingleMax+root.Val,res.SingleMax)
    res.Max=max(left.Max,right.Max)
    res.Max=max(res.Max,root.Val+left.SingleMax+right.SingleMax)
    return res

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
