
// @Title: 最长不含重复字符的子字符串 (最长不含重复字符的子字符串 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 21:18:07
// @Runtime: 8 ms
// @Memory: 3 MB

func lengthOfLongestSubstring(s string) int {
    if len(s)==0{
        return 0
    }

    right:=0
    left:=0
    win:=make(map[byte]int)
    ans:=0
    for right<len(s){
        cur:=s[right]
        right++
        win[cur]++
        
        for win[cur]>1{
            del:=s[left]
            left++
            win[del]--
        }
        ans=max(right-left,ans)
       
    }
    return ans
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
