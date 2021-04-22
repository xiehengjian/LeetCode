
// @Title: 平衡二叉树 (平衡二叉树 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 16:45:15
// @Runtime: 8 ms
// @Memory: 5.7 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root==nil{
        return true
    }
    return math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right)))<=1 && isBalanced(root.Left) && isBalanced(root.Right)

}

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
