
// @Title: 二叉树的深度 (二叉树的深度 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 16:42:06
// @Runtime: 4 ms
// @Memory: 4.2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }
    return max(maxDepth(root.Left),maxDepth(root.Right))+1

}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
