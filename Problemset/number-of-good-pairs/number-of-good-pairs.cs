
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 00:30:34
// @Runtime: 100 ms
// @Memory: 23.5 MB

public class Solution {
    public int NumIdenticalPairs(int[] nums) {
        int count=0;
        for(int i=0;i<nums.Length-1;i++){
            for(int j=i+1;j<nums.Length;j++){
                if(nums[i]==nums[j]){
                    count++;
                }
            }
        }
        return count;

    }
}
