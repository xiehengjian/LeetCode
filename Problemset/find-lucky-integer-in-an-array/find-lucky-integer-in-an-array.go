
// @Title: 找出数组中的幸运数 (Find Lucky Integer in an Array)
// @Author: 846188037@qq.com
// @Date: 2021-03-29 22:36:08
// @Runtime: 4 ms
// @Memory: 3.1 MB

func findLucky(arr []int) int {
    Map:=make(map[int]int)
    for _,v := range arr{
        Map[v]++
    }
    res:=-1
    for k,v := range Map{
        if k==v && v>res{
            res=v
        }
    }
    return res
}
