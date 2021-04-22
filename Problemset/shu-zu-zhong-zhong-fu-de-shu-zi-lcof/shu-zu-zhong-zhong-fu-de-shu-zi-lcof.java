
// @Title: 数组中重复的数字 (数组中重复的数字 LCOF)
// @Author: 846188037@qq.com
// @Date: 2020-12-09 01:56:00
// @Runtime: 0 ms
// @Memory: 46.1 MB

class Solution {
    public int findRepeatNumber(int[] nums) {
        int i = 0;
        while(i < nums.length) {
            if(nums[i] == i) {
                i++;
                continue;
            }
            if(nums[nums[i]] == nums[i]) return nums[i];
            int tmp = nums[i];
            nums[i] = nums[tmp];
            nums[tmp] = tmp;
        }
        return -1;
    }
}


