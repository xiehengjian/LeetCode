
// @Title: 数组中重复的数字 (数组中重复的数字 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 00:14:47
// @Runtime: 48 ms
// @Memory: 9.8 MB

func findRepeatNumber(nums []int) int {
    Map:=make(map[int]int)
    for _,v:=range nums{
        if _,ok:=Map[v];ok{
            return v
        }
        Map[v]++
    }
    return 0
}
