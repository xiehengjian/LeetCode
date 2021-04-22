
// @Title: 同构字符串 (Isomorphic Strings)
// @Author: 846188037@qq.com
// @Date: 2021-01-27 17:45:35
// @Runtime: 4 ms
// @Memory: 2.6 MB

func isIsomorphic(pattern string, s string) bool {
    Map:=make(map[byte]byte)
    strlist:=s
    if len(pattern)!=len(strlist){
        return false
    }
    for i:=0;i<len(pattern);i++{
        v,ok:=Map[pattern[i]]
        if ok && v !=strlist[i] {
            return false
        }
        Map[pattern[i]]=strlist[i]
    }
    check:=make(map[byte]int)
    for _,v :=range Map{
        _,ok:=check[v]
        if ok{
            return false
        }
        check[v]++
    }
    return true
}
