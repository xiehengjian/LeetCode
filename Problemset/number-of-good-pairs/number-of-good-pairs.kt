
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 00:52:26
// @Runtime: 176 ms
// @Memory: 31.3 MB

class Solution {
    fun numIdenticalPairs(nums: IntArray): Int {
        var count=0
        
        for(i in 0..nums.size-2){
            for(j in i+1..nums.size-1){
                if(nums[i]==nums[j]){
                    count++
                }
            }
        }
        return count

    }
}
