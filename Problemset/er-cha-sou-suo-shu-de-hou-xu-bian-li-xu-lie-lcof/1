
// @Title: 二叉搜索树的后序遍历序列 (二叉搜索树的后序遍历序列 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 14:44:49
// @Runtime: 0 ms
// @Memory: 2 MB

func verifyPostorder(postorder []int) bool {
    if len(postorder)<=2{
        return true
    }
    root:=postorder[len(postorder)-1]
    index:=len(postorder)-1
    postorder=postorder[:len(postorder)-1]
    
    for i:=0;i<len(postorder);i++{
        if postorder[i]>root{
            index=i 
            break
        }
    }
    for i:=index;i<len(postorder);i++{
        if postorder[i]<root{
            return false
        }
    }
    return verifyPostorder(postorder[:index])&&verifyPostorder(postorder[index:])

}
