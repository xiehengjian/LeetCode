
// @Title: 一维数组的动态和 (Running Sum of 1d Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-14 21:00:06
// @Runtime: 16 ms
// @Memory: 8.4 MB

class Solution {
public:
    vector<int> runningSum(vector<int>& nums) {
        int sum=0;
        vector<int> runningSum;
        for(int i=0;i<nums.size();i++){
            for(int j=0;j<i+1;j++){
                sum+=nums[j];
            }
            runningSum.push_back(sum);
            sum=0;
        }
        return runningSum;

    }
};
