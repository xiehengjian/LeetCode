
// @Title: 二叉树中和为某一值的路径 (二叉树中和为某一值的路径 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 14:58:02
// @Runtime: 4 ms
// @Memory: 4.6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {

    res:=make([][]int,0)
        if root==nil{
        return res
    }
    temp:=make([]int,0)
    traverse(root,target,&res,temp)
    return res

}

func traverse(root *TreeNode,target int,res *[][]int,temp []int){
    
    if root==nil{
        return 
    }
    
    temp=append(temp,root.Val)
    if root.Left==nil && root.Right==nil && sum(temp)==target{
        fmt.Println(temp)
        ans:=make([]int,len(temp))
        copy(ans,temp)
        *res=append(*res,ans)
        return
    }
    traverse(root.Left,target,res,temp)
    traverse(root.Right,target,res,temp)
}

func sum(a []int)int{
    res:=0
    for _,v:=range a{
        res+=v
    }
    return res
}
