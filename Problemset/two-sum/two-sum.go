
// @Title: 两数之和 (Two Sum)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 18:03:39
// @Runtime: 4 ms
// @Memory: 3 MB

func twoSum(nums []int, target int) []int {
    Map:=make(map[int]int)
    for k,v:=range nums{
        i,ok:=Map[target-v]
        if ok{
            return []int{k,i}
        }
        Map[v]=k
    }
    return []int{}
}
