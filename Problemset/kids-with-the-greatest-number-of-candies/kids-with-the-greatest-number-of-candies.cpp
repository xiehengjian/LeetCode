
// @Title: 拥有最多糖果的孩子 (Kids With the Greatest Number of Candies)
// @Author: 846188037@qq.com
// @Date: 2020-07-16 20:44:04
// @Runtime: 8 ms
// @Memory: 9.2 MB

class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        int n = candies.size();
        int maxCandies = *max_element(candies.begin(), candies.end());
        vector<bool> ret;
        for (int i = 0; i < n; ++i) {
            ret.push_back(candies[i] + extraCandies >= maxCandies);
        }
        return ret;
    }
};


