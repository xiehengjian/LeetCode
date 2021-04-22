
// @Title: 重新排列字符串 (Shuffle String)
// @Author: 846188037@qq.com
// @Date: 2020-08-04 20:11:24
// @Runtime: 4 ms
// @Memory: 3.3 MB

func restoreString(s string, indices []int) string {
    
    ans := make([]byte, len(s))
    for i:=0;i<len(s);i++{
        ans[indices[i]]=s[i]
    }
    return string(ans)

}
