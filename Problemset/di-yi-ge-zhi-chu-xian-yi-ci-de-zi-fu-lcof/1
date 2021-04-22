
// @Title: 第一个只出现一次的字符 (第一个只出现一次的字符  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 16:08:17
// @Runtime: 48 ms
// @Memory: 5.4 MB

func firstUniqChar(s string) byte {
    if s==""{
        return ' '
    }
    Map:=make(map[byte]int)
    for i:=0;i<len(s);i++{
        Map[s[i]]++
    }
    for i:=0;i<len(s);i++{
        if Map[s[i]]==1{
            return s[i]
        }
    }
    return ' '

}
