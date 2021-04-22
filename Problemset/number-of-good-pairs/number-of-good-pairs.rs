
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 01:33:56
// @Runtime: 0 ms
// @Memory: 1.9 MB

impl Solution {
    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        let mut count=0;
        for i in 0..nums.len()-1{
            for j in i+1..nums.len(){
                if nums[i]==nums[j]{
                    count=count+1;
                }
            }
        }
        return count

    }
}
