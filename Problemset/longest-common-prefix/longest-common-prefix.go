
// @Title: 最长公共前缀 (Longest Common Prefix)
// @Author: 846188037@qq.com
// @Date: 2021-03-18 09:59:58
// @Runtime: 0 ms
// @Memory: 2.3 MB

func longestCommonPrefix(strs []string) string {
    if len(strs)==0{
        return ""
    }
    prefix:=strs[0]
    for _,v := range strs{
        for i:=0;i<len(prefix);i++{
            if i==len(v){
                prefix=v
                break
            }
            if prefix[i]!=v[i]{
                prefix=prefix[:i]
                break
            }
        }
    }
    return prefix

}
