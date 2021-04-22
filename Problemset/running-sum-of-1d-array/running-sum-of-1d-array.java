
// @Title: 一维数组的动态和 (Running Sum of 1d Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-14 21:33:14
// @Runtime: 0 ms
// @Memory: 39.2 MB

class Solution {
    public int[] runningSum(int[] nums) {
     for(int i=1;i<nums.length;i++){
         nums[i]+=nums[i-1];
     }
     return nums;
    }
}


