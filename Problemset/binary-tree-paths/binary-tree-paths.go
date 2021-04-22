
// @Title: 二叉树的所有路径 (Binary Tree Paths)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 09:17:21
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
func binaryTreePaths(root *TreeNode) []string {
    res:=[]string{}
    temp:=[]string{}
    traverse(root,&res,temp)
    return res

}

func traverse(root *TreeNode,res *[]string,temp []string){
    if root==nil{
        temp=append(temp,"-1")
        return
    }
    temp=append(temp,strconv.Itoa(root.Val))
    fmt.Println(temp)
    if root.Left==nil&&root.Right==nil{
        *res=append(*res,strings.Join(temp, "->"))
        return
    }
    
    traverse(root.Left,res,temp)
   // temp=temp[:len(temp)-1]
    traverse(root.Right,res,temp)
    //temp=temp[:len(temp)-1]
}
