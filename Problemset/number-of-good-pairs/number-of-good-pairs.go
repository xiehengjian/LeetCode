
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-12 21:12:27
// @Runtime: 0 ms
// @Memory: 2 MB

func numIdenticalPairs(nums []int) int {
    count:=0
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]==nums[j]{
                count++
            }
        }
    }
    return count

}
