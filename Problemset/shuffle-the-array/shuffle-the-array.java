
// @Title: 重新排列数组 (Shuffle the Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-15 20:18:25
// @Runtime: 0 ms
// @Memory: 39.3 MB

class Solution {
    public int[] shuffle(int[] nums, int n) {
        int newNums[] = new int[2*n];
        for(int i=0;i<2*n;i+=2){
            newNums[i]=nums[i/2];
            newNums[i+1]=nums[n+i/2];
        }
        return newNums;

    }
}
