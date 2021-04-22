
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 00:01:59
// @Runtime: 1 ms
// @Memory: 36 MB

class Solution {
    public int numIdenticalPairs(int[] nums) {
        int count=0;
        for(int i=0;i<nums.length-1;i++){
            for(int j=i+1;j<nums.length;j++){
                if(nums[i]==nums[j]){
                    count++;
                }
            }
        }
        return count;

    }
}
