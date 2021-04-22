
// @Title: 独一无二的出现次数 (Unique Number of Occurrences)
// @Author: 846188037@qq.com
// @Date: 2021-01-27 17:22:36
// @Runtime: 0 ms
// @Memory: 2.3 MB

func uniqueOccurrences(arr []int) bool {
    dic1:=make(map[int]int)
    dic2:=make(map[int]int)
    for _,v :=range arr{
        dic1[v]++
    }
    for _,v :=range dic1{
        dic2[v]++
    }
    if len(dic1)==len(dic2){
        return true
    }
    return false
}
