
# @Title: 拥有最多糖果的孩子 (Kids With the Greatest Number of Candies)
# @Author: 846188037@qq.com
# @Date: 2020-07-16 20:11:25
# @Runtime: 48 ms
# @Memory: 13.4 MB

class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        Max=max(candies)
        for i in range(len(candies)):
            if candies[i]+extraCandies>=Max:
                candies[i]=True
            else:
                candies[i]=False
        return candies
