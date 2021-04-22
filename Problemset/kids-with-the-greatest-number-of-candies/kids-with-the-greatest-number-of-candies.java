
// @Title: 拥有最多糖果的孩子 (Kids With the Greatest Number of Candies)
// @Author: 846188037@qq.com
// @Date: 2020-07-16 20:44:53
// @Runtime: 1 ms
// @Memory: 38.8 MB

class Solution {
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int n = candies.length;
        int maxCandies = 0;
        for (int i = 0; i < n; ++i) {
            maxCandies = Math.max(maxCandies, candies[i]);
        }
        List<Boolean> ret = new ArrayList<Boolean>();
        for (int i = 0; i < n; ++i) {
            ret.add(candies[i] + extraCandies >= maxCandies);
        }
        return ret;
    }
}

