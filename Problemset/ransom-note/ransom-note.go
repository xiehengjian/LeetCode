
// @Title: 赎金信 (Ransom Note)
// @Author: 846188037@qq.com
// @Date: 2021-03-29 22:42:03
// @Runtime: 20 ms
// @Memory: 3.8 MB

func canConstruct(ransomNote string, magazine string) bool {
    Map:=make(map[byte]int)
    for i:=0;i<len(magazine);i++{
        Map[magazine[i]]++
    }
    for i:=0;i<len(ransomNote);i++{
        _,ok:=Map[ransomNote[i]]
        if !ok || Map[ransomNote[i]]==0{
            return false
        }else{
            Map[ransomNote[i]]--
        }
    }
    return true
}
