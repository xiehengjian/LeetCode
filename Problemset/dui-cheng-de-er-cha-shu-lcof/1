
// @Title: 对称的二叉树 (对称的二叉树  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:46:14
// @Runtime: 0 ms
// @Memory: 3 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root==nil{
        return true
    }
    
    
    root.Right=mirrorTree(root.Right)
    
    return equal(root.Left,root.Right)
    
}

func equal(A,B *TreeNode)bool{
    fmt.Print(A)
    fmt.Print(" ")
    fmt.Println(B)
    if A==nil && B==nil{
        return true
    }

    if A==nil || B==nil{
        return false
    }
    if A.Val!=B.Val{
        return false
    }
    return equal(A.Left,B.Left) && equal(A.Right,B.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
    if root==nil{
        return root
    }
    temp:=root.Left
    root.Left=root.Right
    root.Right=temp
    root.Left=mirrorTree(root.Left)
    root.Right=mirrorTree(root.Right)
    return root

}
