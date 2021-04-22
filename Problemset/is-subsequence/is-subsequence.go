
// @Title: 判断子序列 (Is Subsequence)
// @Author: 846188037@qq.com
// @Date: 2021-01-27 16:40:57
// @Runtime: 0 ms
// @Memory: 2 MB

func isSubsequence(s string, t string) bool {
/*这个边界其实不要了,在for循环里覆盖了
    if len(s)==0{
        return true
    }
    if len(t)==0{
        return false
    }
    */
        if len(s)>len(t){
        return false
    }
    p1:=0
    p2:=0
    for p1<len(s) && p2<len(t){
        for p1<len(s) && p2<len(t) && s[p1]==t[p2]{//比对成功
            p1++
            p2++
        }
        p2++
    }
    if p1==len(s){//这里要注意下,如果全部匹配了,匹配完最后一个,p1会++,所以p1会越界
        return true
    }
    return false

}
