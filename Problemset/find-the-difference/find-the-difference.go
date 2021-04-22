
// @Title: 找不同 (Find the Difference)
// @Author: 846188037@qq.com
// @Date: 2021-03-18 10:14:10
// @Runtime: 4 ms
// @Memory: 2.2 MB

func findTheDifference(s string, t string) byte {
    Map:=make(map[byte]int)
    for i:=0;i<len(s);i++{
        Map[s[i]]++
    }
    for i:=0;i<len(t);i++{
        if p,ok:=Map[t[i]];!ok || p==0{
            return t[i]
        }
        Map[t[i]]--
    }
    var b byte
    return b
}
