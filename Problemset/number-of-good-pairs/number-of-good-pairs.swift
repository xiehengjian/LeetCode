
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 01:05:13
// @Runtime: 12 ms
// @Memory: 20.2 MB

class Solution {
    func numIdenticalPairs(_ nums: [Int]) -> Int {
        var count=0
        for i in 0...nums.count-2{
            for j in i+1...nums.count-1{
                if nums[i]==nums[j]{
                    count+=1
                }
            }
        }
        return count
    }
}
