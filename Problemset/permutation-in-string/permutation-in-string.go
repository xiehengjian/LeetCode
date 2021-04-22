
// @Title: 字符串的排列 (Permutation in String)
// @Author: 846188037@qq.com
// @Date: 2021-01-17 15:37:03
// @Runtime: 20 ms
// @Memory: 2.6 MB

func checkInclusion(s1 string, s2 string) bool {
    win:=make(map[byte]int)
    need:=make(map[byte]int)
    for i:=0;i<len(s1);i++{
        need[s1[i]]++
    }
    left:=0
    right:=0
    match:=0
    for right<len(s2){
        c:=s2[right]
        right++
        if need[c]!=0{
            win[c]++
            if win[c]==need[c]{
                match++
            }
        }
        for right-left>=len(s1){
            if match == len(need){
                return true
            }
            d:=s2[left]
            left++
            if need[d]!=0{
                if win[d]==need[d]{
                    match--
                }
                win[d]--
            }
        }
    }
    return false

}
