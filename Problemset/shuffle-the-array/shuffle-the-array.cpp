
// @Title: 重新排列数组 (Shuffle the Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-15 20:22:34
// @Runtime: 8 ms
// @Memory: 9.6 MB

class Solution {
public:
    vector<int> shuffle(vector<int>& nums, int n) {
        vector<int> newNums;
        for(int i=0;i<n;i++){
            newNums.push_back(nums[i]);
            newNums.push_back(nums[i+n]);

        }
        return newNums;

    }
};
