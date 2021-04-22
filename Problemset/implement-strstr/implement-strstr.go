
// @Title: 实现 strStr() (Implement strStr())
// @Author: 846188037@qq.com
// @Date: 2020-09-28 13:18:42
// @Runtime: 0 ms
// @Memory: 2.2 MB

func strStr(haystack string, needle string) int {
    if len(needle)==0{
        return 0
    }
    var i,j int
    for i=0;i<len(haystack)-len(needle)+1;i++{
        for j=0;j<len(needle);j++{
            if haystack[i+j]!=needle[j]{
                break
            }
        }
        if len(needle)==j{
            return i
        }
    }
    return -1

}    
