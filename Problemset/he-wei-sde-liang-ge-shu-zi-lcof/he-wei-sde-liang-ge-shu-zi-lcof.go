
// @Title: 和为s的两个数字 (和为s的两个数字 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 17:23:45
// @Runtime: 300 ms
// @Memory: 11.5 MB

func twoSum(nums []int, target int) []int {
    Map:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        Map[nums[i]]++
    }
    for i:=0;i<len(nums);i++{
        if _,ok:=Map[target-nums[i]];ok{
            return []int{nums[i],target-nums[i]}
        }
    }
    return []int{}

}
