
// @Title: 最小覆盖子串 (Minimum Window Substring)
// @Author: 846188037@qq.com
// @Date: 2021-04-21 17:08:53
// @Runtime: 156 ms
// @Memory: 2.9 MB

func minWindow(s string, t string) string {
    left,right:=0,0
    win:=make(map[byte]int)
    sub:=make(map[byte]int)
    for i:=0;i<len(t);i++{
        sub[t[i]]++
    }
    ans:=s
    for right<len(s){
        for right<len(s) && !isCover(win,sub){
            if len(win)==len(s){
                return ""
            }
            win[s[right]]++
            right++   
        }
        for left<=right && isCover(win,sub){
            if len(ans)>=right-left{
                ans=s[left:right]                
            }
            win[s[left]]-- 
            left++ 
        }
    }
    if left==0{
        return ""
    }
    return ans
}



func isCover(win,sub map[byte]int)bool{
    for k,v:=range sub{
        p,ok:=win[k]
        if !ok || p<v{
            return false
        }
    }
    return true
}
